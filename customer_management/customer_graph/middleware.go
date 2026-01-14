package graph

import (
	"credit_system/auth_service/helpers"
	"credit_system/auth_service/model"
	"credit_system/auth_service/services"

	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
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
			owner, _ = userInterface.(*model.StoreOwner)
		}

		if owner == nil {
			return helpers.FormatError(fmt.Errorf("invalid_token"))
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
