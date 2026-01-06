package repositories

import (
	"auth_service/model"

	"context"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {

	//User
	CheckForExistingUser(field, value string) (*model.User, error)
	FetchUserByLoginID(field, value string) (*model.User, error)
	RegisterUser(signUpInput *model.SignupInput) (*model.User, *model.Token, *model.AuthUserProfile, error)
	Login(loginId, password string) (*model.User, *model.Token, error)
	UpdateSingleDataByID(userID uuid.UUID, field, value string) (*model.User, error)
	FetchUser(userID uuid.UUID) (*model.User, error)

	// Token & Authentication
	CreateToken(user *model.User, clientId, clientSecret string, tx *gorm.DB) (*model.Token, error)
	CreateRefreshToken(tokenId string) (*model.RefreshToken, error)
	FindAccessToken(id uuid.UUID) (*model.AccessToken, error)
	FindRefreshToken(id uuid.UUID) (*model.RefreshToken, error)
	FindRefreshTokenByAccessToken(accessToken string) (*model.RefreshToken, error)
	FindAccessTokenForUser(id uuid.UUID, userId uuid.UUID) (*model.AccessToken, error)
	FindAccessTokensForUser(userId uuid.UUID) ([]model.AccessToken, error)
	GetValidAccessToken(userId uuid.UUID, clientId string) (*model.AccessToken, error)
	GetLatestValidAccessToken(userId uuid.UUID) (*model.AccessToken, error)
	RevokeAccessToken(id uuid.UUID) (*model.AccessToken, error)
	RevokeRefreshToken(id uuid.UUID) (*model.RefreshToken, error)
	DeleteTokens(userId uuid.UUID) error
	DeleteOldestAccessToken(userId uuid.UUID) error
	GetOldestAccessToken(userId uuid.UUID) (*model.AccessToken, error)
	DeleteRefreshToken(accessTokenId string) error
	DeleteAccessToken(id string) error
	VerifyClient(clientId, secretKey string) bool
	FindAccessTokenByIdentifier(identifier string) (*model.AccessToken, error)
	FindRefreshTokenByIdentifier(identifier string) (*model.RefreshToken, error)

	CreateUserActivity(ctx context.Context, inputActivityType string, userID uuid.UUID) (*model.UserActivity, error)
	CreateGameActivity(ctx context.Context, userID uuid.UUID) error
	DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error)
}
