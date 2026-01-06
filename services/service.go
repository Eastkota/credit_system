package services

import (
	"auth_service/model"
	"context"

	"github.com/google/uuid"
)

type Services interface {
	CheckForExistingUser(field, value string) (*model.User, error)
	ValidateToken(tokenString string) (*model.User, error)
	RefreshToken(tokenString string) (*model.User, *model.Token, *model.AuthUserProfile, *model.AuthUserMembership, error)
	IsValidToken(tokenString string) (string, bool)
	Signup(signupData model.SignupInput) (*model.User, *model.Token, *model.AuthUserProfile, error)
	Login(loginId, password string) (*model.User, *model.Token, *model.AuthUserProfile, *model.AuthUserMembership, error)
	Logout(tokenId string) error
	UpdatePassword(updatePaswordData model.UpdatePasswordInput) error
	UpdateSingleDataByID(userId uuid.UUID, field, value, password string) (*model.User, error)
	FetchUser(userId uuid.UUID) (*model.User, error)
	ResetPassword(userId uuid.UUID, password, confirmPassword string) (error)

	SaveUserActivity(ctx context.Context, input *model.UserActivityInput) (*model.UserActivity, error)
	SaveGameActivity(ctx context.Context, userID uuid.UUID) error
	DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error)
}
