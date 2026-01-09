package services

import (
	"credit_system/auth_service/model"

	"github.com/google/uuid"
)

type Services interface {
	CheckForExistingUser(field, value string) (*model.StoreOwner, error)
	ValidateToken(tokenString string) (*model.StoreOwner, error)
	RefreshToken(tokenString string) (*model.StoreOwner, *model.Token, *model.Store, error)
	IsValidToken(tokenString string) (string, bool)
	Signup(signupData model.SignupInput) (*model.StoreOwner, *model.Token, *model.Store, error)
	Login(phoneNumber, password string) (*model.StoreOwner, *model.Token, *model.Store, error)
	// Logout(tokenId string) error
	// UpdatePassword(updatePaswordData model.UpdatePasswordInput) error
	UpdateSingleDataByID(ownerID uuid.UUID, field, value, password string) (*model.StoreOwner, error)
	FetchOwnerByID(ownerID uuid.UUID) (*model.StoreOwner, error)
	FetchStore(storeID uuid.UUID) (*model.Store, error)
	FetchStoreByOwnerID(ownerID uuid.UUID) (*model.Store, error)
	// ResetPassword(userId uuid.UUID, password, confirmPassword string) (error)

	// SaveUserActivity(ctx context.Context, input *model.UserActivityInput) (*model.UserActivity, error)
	// SaveGameActivity(ctx context.Context, userID uuid.UUID) error
	// DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error)
}
