package services

import (
	"credit_system/config"
	"credit_system/auth_service/model"
	"credit_system/auth_service/repositories"

	"fmt"
	"time"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
    ErrInvalidToken    = fmt.Errorf("invalid_token")
    ErrRefreshRequired = fmt.Errorf("refresh_required")
)

type AuthService struct {
	Repository repositories.Repository // Inject Repository
}

func NewService(repository repositories.Repository) *AuthService {
	return &AuthService{Repository: repository}
}

func (as *AuthService) CheckForExistingUser(field, value string) (*model.StoreOwner, error) {
	return as.Repository.CheckForExistingUser(field, value)
}

func (as *AuthService) Signup(signupData model.SignupInput) (*model.StoreOwner, *model.Token, *model.Store, error) {

	if signupData.PhoneNumber == "" {
		return nil, nil, nil, fmt.Errorf("either email or phone number is required")
	}
	if signupData.PhoneNumber != "" {
		mobileResult, err := as.CheckForExistingUser("phone_number", signupData.PhoneNumber)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to register: %v", err)
		}

		if mobileResult != nil {
			return nil, nil, nil, fmt.Errorf("user with given mobile number %v already exist", signupData.PhoneNumber)
		}
	}
	return as.Repository.RegisterUser(&signupData)
}

func (as *AuthService) FetchStore(storeID uuid.UUID) (*model.Store, error) {
	return as.Repository.FetchStore(storeID)
}

func (as *AuthService) FetchStoreByOwnerID(ownerID uuid.UUID) (*model.Store, error) {
	return as.Repository.FetchStoreByOwnerID(ownerID)
}

func (as *AuthService) Login(phoneNumber, password string) (*model.StoreOwner, *model.Token, *model.Store, error) {
	owner, token, store, err := as.Repository.Login(phoneNumber, password)
	if err != nil {
		return nil, nil, nil, err
	}

	store, err = as.Repository.FetchStore(store.ID)
	if err != nil {
		return nil, nil, nil, err
	}

	return owner, token, store, nil
}

// func (as *AuthService) Logout(tokenId string) error {
// 	if tokenId == "" {
// 		return fmt.Errorf("invalid_token")
// 	}
// 	err := as.Repository.DeleteAccessToken(tokenId)
// 	if err != nil {
// 		return fmt.Errorf("failed to logout: %v", err)
// 	}
// 	return nil
// }

func (as *AuthService) ValidateToken(authHeader string) (*model.StoreOwner, error) {
    if strings.TrimSpace(authHeader) == "" {
        return nil, ErrInvalidToken
    }

    tokenStr := strings.TrimSpace(authHeader)
    if len(tokenStr) > 6 && strings.EqualFold(tokenStr[:7], "bearer ") {
        tokenStr = strings.TrimSpace(tokenStr[7:])
    }
    if tokenStr == "" {
        return nil, ErrInvalidToken
    }

    claims := &model.Claims{}
    parsedToken, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return []byte(config.ClientSecret()), nil
    })
    if err != nil {
        return nil, ErrInvalidToken
    }
    if !parsedToken.Valid {
        return nil, ErrInvalidToken
    }

    if claims.TokenId == "" {
        return nil, ErrInvalidToken
    }

    accessToken, err := as.Repository.FindAccessTokenByPhoneNumber(claims.TokenId)
    if err != nil || accessToken == nil {
        return nil,  ErrInvalidToken
    }

    if accessToken.Revoked {
        return nil, ErrInvalidToken
    }

    now := time.Now()

    if accessToken.ExpiresAt.After(now) {
        var owner *model.StoreOwner
        if claims.Owner != nil && claims.Owner.ID != uuid.Nil {
            owner = claims.Owner
        } else if claims.Owner == nil || claims.Owner.ID == uuid.Nil {
            return nil,  ErrInvalidToken
        } else {
            u, ferr := as.Repository.FetchOwnerByID(claims.Owner.ID) 
            if ferr != nil {
                return nil,  ErrInvalidToken
            }
            owner = u
        }
        return owner, nil
    }

    refreshRec, err := as.Repository.FindRefreshTokenByAccessToken(accessToken.Token)
    if err != nil || refreshRec == nil {
        return nil, ErrInvalidToken
    }

    if refreshRec.Revoked || !refreshRec.ExpiresAt.Before(now) {
        return nil, ErrInvalidToken
    }
	return nil, ErrRefreshRequired
}

func (as *AuthService) RefreshToken(refreshTokenStr string) (*model.StoreOwner, *model.Token, *model.Store, error) {
	if strings.TrimSpace(refreshTokenStr) == "" {
        return nil, nil, nil, ErrInvalidToken
    }

    claims := &model.Claims{}
    parsedToken, err := jwt.ParseWithClaims(refreshTokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return []byte(config.ClientSecret()), nil
    })
    if err != nil {
        return nil, nil, nil, ErrInvalidToken
    }
    if !parsedToken.Valid {
        return nil, nil, nil, ErrInvalidToken
    }

    if claims.TokenId == "" {
        return nil, nil, nil, ErrInvalidToken
    }

    refreshToken, err := as.Repository.FindRefreshTokenByPhoneNumber(claims.TokenId)
    if err != nil || refreshToken == nil {
        return nil, nil, nil, ErrInvalidToken
    }

    if refreshToken.Revoked {
        return nil, nil, nil, ErrInvalidToken
    }

    if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, nil, nil, ErrInvalidToken
	}else{
        var owner *model.StoreOwner
        if claims.Owner == nil || claims.Owner.ID == uuid.Nil {
            return nil, nil, nil, ErrInvalidToken
        } else {
            u, ferr := as.Repository.FetchOwnerByID(claims.Owner.ID) 
            if ferr != nil {
                return nil, nil, nil, ErrInvalidToken
            }
            owner = u
			token, err := as.Repository.CreateToken(owner, config.ClientId(), config.ClientSecret(), nil)
			if err != nil {
				return nil, nil, nil, ErrInvalidToken
			}
			store, err := as.Repository.FetchStore(owner.ID)
			if err != nil {
				return nil, nil, nil, ErrInvalidToken
			}
			return owner, token, store, nil
        }
        return nil, nil, nil, ErrInvalidToken
    }
	return nil, nil, nil, ErrInvalidToken
}

func (as *AuthService) IsValidToken(tokenString string) (string, bool) {
	if tokenString != "" {
		var clientSecret = []byte(config.ClientSecret())
		token, err := jwt.ParseWithClaims(tokenString[7:], &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return clientSecret, nil
		})
		if err != nil {
			return "",false
		}
		claims, ok := token.Claims.(*model.Claims)
		if ok {
			accessToken, err := as.Repository.FindAccessTokenByPhoneNumber(claims.TokenId)

			if err != nil || accessToken == nil {
				return "",false
			}

			if accessToken.Revoked {
				return "",false
			}

			if !accessToken.ExpiresAt.After(time.Now()) {
				return "",false
			}
			return claims.TokenId,true
		}
	}
	return "", false
}

// func (as *AuthService) UpdatePassword(updatePaswordData model.UpdatePasswordInput) error {
// 	if updatePaswordData.CurrentPassword == "" {
// 		return fmt.Errorf("current password cannot be empty")
// 	}
// 	if updatePaswordData.Password == "" || updatePaswordData.ConfirmPassword == "" {
// 		return fmt.Errorf("new password and confirm password cannot be empty")
// 	}
// 	if updatePaswordData.Password != updatePaswordData.ConfirmPassword {
// 		return fmt.Errorf("new password and confirm password doesnot match")
// 	}

// 	user, err := as.Repository.FetchUser(updatePaswordData.UserId)
// 	if err != nil {
// 		return err
// 	}

// 	if !helpers.IsValidPassword(updatePaswordData.CurrentPassword, user.Password) {
// 		return fmt.Errorf("current doesnot match")
// 	}

// 	password, err := helpers.EncryptPassword(updatePaswordData.Password)

// 	if err != nil {
// 		return err
// 	}

// 	_, err = as.Repository.UpdateSingleDataByID(updatePaswordData.UserId, "password", password)

// 	if err != nil {
// 		return fmt.Errorf("failed to update password: %v", err)
// 	}

// 	return nil
// }

// func (as *AuthService) UpdateSingleDataByID(ownerID uuid.UUID, field, value, password string) (*model.StoreOwner, error) {
// 	owner, err := as.Repository.FetchUser(OwnerID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !helpers.IsValidPassword(password, owner.Password) {
// 		return nil, fmt.Errorf("authentication failed with provided password")
// 	}

// 	return as.Repository.UpdateSingleDataByID(ownerID, field, value)
// }

func (as *AuthService) FetchOwnerByID(ownerID uuid.UUID) (*model.StoreOwner, error) {
	return as.Repository.FetchOwnerByID(ownerID)
}

// func (as *AuthService) ResetPassword(userID uuid.UUID, password, confirmPassword string) error {
// 	if password == "" || confirmPassword == "" {
// 		return fmt.Errorf("new password and confirm password cannot be empty")
// 	}
// 	if password != confirmPassword {
// 		return fmt.Errorf("new password and confirm password doesnot match")
// 	}

// 	newPassword, err := helpers.EncryptPassword(password)

// 	if err != nil {
// 		return err
// 	}

// 	_, err = as.Repository.UpdateSingleDataByID(userID, "password", newPassword)

// 	if err != nil {
// 		return fmt.Errorf("failed to update password: %v", err)
// 	}

// 	return nil
// }

// func (as *AuthService) DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error) {
// 	return as.Repository.DeleteUser(ctx, userID, status, password)
// }


// func (as *AuthService) SaveUserActivity(ctx context.Context, input *model.UserActivityInput) (*model.UserActivity, error) {
//     loadedActivity, err := as.Repository.CreateUserActivity(ctx, input.Activity, input.UserID)
//     if err != nil {
//         return nil, err
//     }

//     return loadedActivity, nil
// }

// func (as *AuthService) SaveGameActivity(ctx context.Context, userID uuid.UUID) error {
//     err := as.Repository.CreateGameActivity(ctx, userID)
//     if err != nil {
//         return err
//     }

//     return nil
// }