package schema

import (
    "credit_system/auth_service/helpers"
    "credit_system/auth_service/model"
    "credit_system/auth_service/services"
    "credit_system/core/generic_response"

    "fmt"
    "context"
    "net/http"
    "os"
    "time"
    "crypto/sha256"

    "github.com/graphql-go/graphql"
    // "github.com/google/uuid"
)
var authService *services.AuthService 

func InitMiddleware(as *services.AuthService) {
    authService = as
}

func AuthMiddleware[T any](next func(p graphql.ResolveParams) *T) func(p graphql.ResolveParams) *T {
    return func(p graphql.ResolveParams) *T {
        ctx := p.Context
        userInterface := ctx.Value(model.UserKey)

        if userInterface == nil {
            if req, ok := ctx.Value(model.RequestKey).(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := authService.ValidateToken(authHeader)
                
                if err != nil {
                    // This will now return the correct struct type automatically
                    return generic_response.FormatMiddlewareError[T](err)
                }

                ctx = context.WithValue(ctx, model.UserKey, u)
                p.Context = ctx
            } else {
                return generic_response.FormatMiddlewareError[T](fmt.Errorf("invalid_request_context"))
            }
        }

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