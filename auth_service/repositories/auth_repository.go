package repositories

import (
    "auth_service/model"
    "auth_service/config"
    "auth_service/helpers"

    "fmt"
    "time"

    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type AuthRepository struct {
    DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
    return &AuthRepository{DB: db}
}

func (repo *AuthRepository) CreateToken(owner *model.StoreOwner, clientId, clientSecret string, tx *gorm.DB) (*model.Token, error) {
	if !repo.VerifyClient(clientId, clientSecret) {	
		return nil, fmt.Errorf("invalid client credentials")
	}
	if owner == nil {
		return nil, fmt.Errorf("owner cannot be nil")
	}
	accessTokenId := uuid.New()
	accessToken, err := helpers.GenerateRandomTokenString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token string: %v", err)
	}
	parsedClientID, err := uuid.Parse(clientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client ID format: %w", err)
	}
	newAccessToken := model.AccessToken{
		ID:             accessTokenId,
		Token:       	accessToken,
		OwnerId:     	owner.ID,
		ClientId:   	parsedClientID,
		Revoked:     	false,
		ExpiresAt:   	time.Now().Add(config.AccessTokenDuration),
		CreatedAt:   	time.Now(),
		UpdatedAt:   	time.Now(),
	}
    if tx == nil {
        tx = repo.DB
    }
	err = tx.Create(&newAccessToken).Error
	if  err != nil {
        tx.Rollback()
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

    refreshTokenId := uuid.New()
    refreshToken, err := helpers.GenerateRandomTokenString(32)
    if err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("failed to generate token string: %v", err)
    }
    newRefreshToken := model.RefreshToken{
        ID:          refreshTokenId,
        Token:       refreshToken,
        AccessToken: accessToken,
        Revoked:     false,
        ExpiresAt:   time.Now().Add(config.RefreshTokenDuration),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    if err := tx.Create(&newRefreshToken).Error; err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("failed to create refresh token: %v", err)
    }

    accessTokenString,err := helpers.GenerateJwtToken(accessToken, config.AccessTokenDuration, clientSecret, owner)
    if err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("failed to generate access token string: %v", err)
    }

    refreshTokenString,err := helpers.GenerateJwtToken(refreshToken, config.RefreshTokenDuration, clientSecret, owner)
    if err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("failed to generate access token string: %v", err)
    }

	tokenResult := &model.Token{
		TokenType:    "Bearer",
		AccessToken:  accessTokenString,
        RefreshToken: refreshTokenString,
		ExpiresIn:    newAccessToken.ExpiresAt,
	}
	return tokenResult, nil
}

func (repo *AuthRepository) CreateRefreshToken(accessToken string) (*model.RefreshToken, error) {
    refreshTokenId := uuid.New()
    token, err := helpers.GenerateRandomTokenString(32)
    if err != nil {
        return nil, fmt.Errorf("failed to generate token string: %v", err)
    }
    refreshToken := model.RefreshToken{
        ID:          refreshTokenId,
        Token:       token,
        AccessToken: accessToken,
        Revoked:     false,
        ExpiresAt:   time.Now().Add(config.RefreshTokenDuration),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    if err := repo.DB.Create(&refreshToken).Error; err != nil {
        return nil, fmt.Errorf("failed to create refresh token: %v", err)
    }
    return &refreshToken, nil
}

func (repo *AuthRepository) FindAccessToken(id uuid.UUID) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.First(&token, "id = ?", id).Error
    if err != nil {
        return nil, fmt.Errorf("access token not found: %v", err)
    }
    return &token, nil
}

func (repo *AuthRepository) FindAccessTokenByPhoneNumber(phone_no string) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.First(&token, "token = ?", phone_no).Error
    if err != nil {
        return nil, fmt.Errorf("access token not found: %v", err)
    }
    return &token, nil
}
func (repo *AuthRepository) FindRefreshTokenByPhoneNumber(phone_no string) (*model.RefreshToken, error) {
    var token model.RefreshToken
    err := repo.DB.First(&token, "token = ?", phone_no).Error
    if err != nil {
        return nil, fmt.Errorf("refresh token not found: %v", err)
    }
    return &token, nil
}

func (repo *AuthRepository) FindRefreshToken(id uuid.UUID) (*model.RefreshToken, error) {
    var token model.RefreshToken
    err := repo.DB.First(&token, "id = ?", id).Error
    if err != nil {
        return nil, fmt.Errorf("refresh token not found: %v", err)
    }
    return &token, nil
}
func (repo *AuthRepository) FindRefreshTokenByAccessToken(accessToken string) (*model.RefreshToken, error) {
    var token model.RefreshToken
    err := repo.DB.First(&token, "access_token = ?", accessToken).Error
    if err != nil {
        return nil, fmt.Errorf("refresh token not found: %v", err)
    }
    return &token, nil
}

func (repo *AuthRepository) FindAccessTokenForUser(id, ownerId uuid.UUID) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.Where("id = ? AND owner_id = ?", id, ownerId).First(&token).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) FindAccessTokensForUser(ownerId uuid.UUID) ([]model.AccessToken, error) {
    var tokens []model.AccessToken
    err := repo.DB.Where("owner_id = ?", ownerId).Find(&tokens).Error
    if err != nil {
        return nil, fmt.Errorf("failed to fetch tokens: %v", err)
    }
    return tokens, nil
}

func (repo *AuthRepository) GetValidAccessToken(ownerId uuid.UUID, clientId string) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.Where("owner_id = ? AND client_id = ? AND revoked = ? AND expires_at > ?", ownerId, clientId, false, time.Now()).First(&token).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) GetLatestValidAccessToken(userId uuid.UUID) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.Where("user_id = ? AND revoked = ? AND expires_at > ?", userId, false, time.Now()).Order("expires_at desc").First(&token).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) RevokeAccessToken(id uuid.UUID) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.Model(&token).Where("id = ?", id).Update("revoked", true).Error
    if err != nil {
        return nil, fmt.Errorf("failed to revoke access token: %v", err)
    }
    err = repo.DB.First(&token, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) RevokeRefreshToken(id uuid.UUID) (*model.RefreshToken, error) {
    var token model.RefreshToken
    err := repo.DB.Model(&token).Where("id = ?", id).Update("revoked", true).Error
    if err != nil {
        return nil, fmt.Errorf("failed to revoke refresh token: %v", err)
    }
    err = repo.DB.First(&token, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) DeleteAccessToken(token string) error {
    if err := repo.DeleteRefreshToken(token); err != nil {
        return fmt.Errorf("failed to delete refresh tokens: %v", err)
    }

    if err := repo.DB.Delete(&model.AccessToken{}, "token = ?", token).Error; err != nil {
        return fmt.Errorf("failed to delete access token: %v", err)
    }
    return nil
}

func (repo *AuthRepository) DeleteRefreshToken(accessTokenId string) error {
    if err := repo.DB.Where("access_token = ?", accessTokenId).Delete(&model.RefreshToken{}).Error; err != nil {
        return fmt.Errorf("failed to delete refresh tokens: %v", err)
    }
    return nil
}

func (repo *AuthRepository) DeleteTokens(ownerId uuid.UUID) error {
    var tokens []model.AccessToken
    if err := repo.DB.Where("owner_id = ?", ownerId).Find(&tokens).Error; err != nil {
        return err
    }
    return repo.DB.Transaction(func(tx *gorm.DB) error {
        for _, accessToken := range tokens {
            if err := tx.Where("access_token = ?", accessToken.ID).Delete(&model.RefreshToken{}).Error; err != nil {
                return err
            }
        }
        if err := tx.Where("owner_id = ?", ownerId).Delete(&model.AccessToken{}).Error; err != nil {
            return err
        }
        return nil
    })
}

func (repo *AuthRepository) GetOldestAccessToken(ownerId uuid.UUID) (*model.AccessToken, error) {
    var token model.AccessToken
    err := repo.DB.Where("owner_id = ? AND revoked = ? AND expires_at > ?", ownerId, false, time.Now()).Order("expires_at asc").First(&token).Error
    if err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *AuthRepository) DeleteOldestAccessToken(ownerId uuid.UUID) error {
	oldestToken, err := repo.GetOldestAccessToken(ownerId)
	if err != nil {
		return fmt.Errorf("failed to get oldest access token: %v", err)
	}
	return repo.DeleteAccessToken(oldestToken.Token)
}

func (repo *AuthRepository) VerifyClient(clientId, secretKey string) bool {
    var client model.Client
    err := repo.DB.First(&client, "id = ?", clientId).Error
    if err != nil {
        return false
    }
    err = bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secretKey))
    return err == nil
}
