package schema

import (
    "auth_service/helpers"
    "auth_service/model"
    "auth_service/services"

    "fmt"
    "context"
    "net/http"
    "os"
    "time"
    "crypto/sha256"
    "log"

    "github.com/graphql-go/graphql"
    // "github.com/google/uuid"
)
var authService *services.AuthService 

func InitMiddleware(as *services.AuthService) {
    authService = as
}

func AuthMiddleware(next func(p graphql.ResolveParams) *model.GenericAuthResponse) func(p graphql.ResolveParams) *model.GenericAuthResponse {
    return func(p graphql.ResolveParams) *model.GenericAuthResponse {
        ctx := p.Context
        userInterface := ctx.Value(model.UserKey)

        var owner *model.StoreOwner
        if userInterface == nil {
            if req, ok := ctx.Value(model.RequestKey).(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := authService.ValidateToken(authHeader)
                if err != nil {
                    return helpers.FormatError(err)
                }
                if u == nil {
                    return helpers.FormatError(fmt.Errorf("invalid_token"))
                }

                ctx = context.WithValue(ctx, model.UserKey, u)
                p.Context = ctx
                owner = u
            } else {
                return helpers.FormatError(fmt.Errorf("invalid_token"))
            }
        } else {
            owner, _ = userInterface.(*model.User)
        }

        if owner == nil {
            return helpers.FormatError(fmt.Errorf("invalid_token"))
        }

        go func() {
            ctx := context.Background() // new context not tied to request
            _, err := authService.SaveUserActivity(ctx, &model.UserActivityInput{
                Activity: p.Info.FieldName,
                ownerID:   owner.ID,
            })
            if err != nil {
                log.Printf("[UserActivity] Failed to save activity: %v", err)
            }
        }()

        return next(p)
    }
}


func PublicAuthMiddleware(next func(p graphql.ResolveParams) *model.GenericAuthResponse) func(p graphql.ResolveParams) *model.GenericAuthResponse {
    return func(p graphql.ResolveParams) *model.GenericAuthResponse {
        ctx := p.Context
        req, ok := ctx.Value(model.RequestKey).(*http.Request)
        if !ok {
            return helpers.FormatError(fmt.Errorf("invalid_request"))
        }

        authHeader := req.Header.Get("Authorization")
        if authHeader == "" {
            return helpers.FormatError(fmt.Errorf("UnAuthorized"))
        }

        envPublicToken := os.Getenv("PUBLIC_ACCESS_TOKEN")
        dailyToken := generateYearlyPublicToken() // optional rotating token

        if authHeader == envPublicToken || authHeader == fmt.Sprintf("Bearer %s", envPublicToken) ||
            authHeader == dailyToken || authHeader == fmt.Sprintf("Bearer %s", dailyToken) {

            ctx = context.WithValue(ctx, "isPublic", true)
            p.Context = ctx
            return next(p)
        }

        u, err := authService.ValidateToken(authHeader)
        if err != nil {
            return helpers.FormatError(err)
        }
        if u == nil {
            return helpers.FormatError(fmt.Errorf("invalid_token"))
        }

        ctx = context.WithValue(ctx, model.UserKey, u)
        p.Context = ctx

        return next(p)
    }
}

func generateYearlyPublicToken() string {
    secret := os.Getenv("PUBLIC_TOKEN_SECRET")

    currentYear := time.Now().Format("2006")

    hash := sha256.Sum256([]byte(secret + currentYear))
    return fmt.Sprintf("%x", hash)
}