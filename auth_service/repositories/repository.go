package repositories

import (
	"credit_system/auth_service/model"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {

	//User
	CheckForExistingUser(field, value string) (*model.StoreOwner, error)
	FetchOwner(field, value string) (*model.StoreOwner, error)
	RegisterUser(signupInput *model.SignupInput) (*model.StoreOwner, *model.Token, *model.Store, error)
	Login(PhoneNumber, password string) (*model.StoreOwner, *model.Token, *model.Store, error)
	UpdateSingleDataByID(ownerID uuid.UUID, field, value string) (*model.StoreOwner, error)
	FetchOwnerByID(ownerID uuid.UUID) (*model.StoreOwner, error)
	FetchStore(storeID uuid.UUID) (*model.Store, error)
	FetchStoreByOwnerID(ownerID uuid.UUID) (*model.Store, error)

	// Token & Authentication
	CreateToken(owner *model.StoreOwner, clientId, clientSecret string, tx *gorm.DB) (*model.Token, error)
	CreateRefreshToken(tokenId string) (*model.RefreshToken, error)
	FindAccessToken(id uuid.UUID) (*model.AccessToken, error)
	FindRefreshToken(id uuid.UUID) (*model.RefreshToken, error)
	FindRefreshTokenByAccessToken(accessToken string) (*model.RefreshToken, error)
	FindAccessTokenForUser(id uuid.UUID, ownerId uuid.UUID) (*model.AccessToken, error)
	FindAccessTokensForUser(ownerId uuid.UUID) ([]model.AccessToken, error)
	GetValidAccessToken(ownerId uuid.UUID, clientId string) (*model.AccessToken, error)
	GetLatestValidAccessToken(ownerId uuid.UUID) (*model.AccessToken, error)
	RevokeAccessToken(id uuid.UUID) (*model.AccessToken, error)
	RevokeRefreshToken(id uuid.UUID) (*model.RefreshToken, error)
	DeleteTokens(ownerId uuid.UUID) error
	DeleteOldestAccessToken(ownerId uuid.UUID) error
	GetOldestAccessToken(ownerId uuid.UUID) (*model.AccessToken, error)
	DeleteRefreshToken(accessTokenId string) error
	DeleteAccessToken(id string) error
	VerifyClient(clientId, secretKey string) bool
	FindAccessTokenByPhoneNumber(phone_number string) (*model.AccessToken, error)
	FindRefreshTokenByPhoneNumber(phone_number string) (*model.RefreshToken, error)

	// CreateUserActivity(ctx context.Context, inputActivityType string, userID uuid.UUID) (*model.UserActivity, error)
	// CreateGameActivity(ctx context.Context, userID uuid.UUID) error
	// DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error)
}
