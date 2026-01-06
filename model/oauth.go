package model

import (
    "time"

    "github.com/google/uuid"
    "github.com/golang-jwt/jwt"
)

type ContextKey string

const (
    TokenKey   ContextKey = "token_id"
    UserKey    ContextKey = "user"
    RequestKey ContextKey = "http_request"
)

type AccessToken struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    UserId    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    ClientId  uuid.UUID `gorm:"type:uuid;not null" json:"client_id"`
    Token     string    `gorm:"type:text;not null" json:"token"`
    Scopes    []string  `gorm:"type:text[]" json:"scopes"`
    Revoked   bool      `gorm:"type:boolean;default:false" json:"revoked"`
    ExpiresAt time.Time `gorm:"type:timestamptz;not null" json:"expires_at"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
func (AccessToken) TableName() string {
    return "auth.access_tokens"
}

type Client struct {
    ID                   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name                 string    `gorm:"type:varchar(255);not null" json:"name"`
    Secret               string    `gorm:"type:text;not null" json:"secret"`
    Provider             string    `gorm:"type:varchar(100)" json:"provider"`
    Redirect             string    `gorm:"type:varchar(255)" json:"redirect"`
    PersonalAccessClient bool      `gorm:"type:boolean;default:false" json:"personal_access_client"`
    PasswordClient       bool      `gorm:"type:boolean;default:false" json:"password_client"`
    Revoked              bool      `gorm:"type:boolean;default:false" json:"revoked"`
    CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
func (Client) TableName() string {
    return "auth.auth_clients"
}

type RefreshToken struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    AccessToken string   `gorm:"type:text;not null" json:"access_token"`
    Token      string    `gorm:"type:text;not null" json:"token"`
    Revoked    bool      `gorm:"type:boolean;default:false" json:"revoked"`
    ExpiresAt  time.Time `gorm:"type:timestamptz;not null" json:"expires_at"`
    CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (RefreshToken) TableName() string {
    return "auth.refresh_tokens"
}

type Token struct {
    TokenType    string `json:"token_type"`
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    time.Time  `json:"expires_in"` // seconds until expiry
}

type Claims struct {
    User    *User  `json:"user"`
    TokenId string `json:"token_id"`
    Exp     string  `json:"exp"`
    Issuer string  `json:"issuer"`
    jwt.StandardClaims
}