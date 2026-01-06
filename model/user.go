package model

import (
    "time"

    "github.com/google/uuid"
)


type User struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    UserIdentifier string    `gorm:"type:varchar(32);unique;not null" json:"user_identifier"`
    Email          string    `gorm:"type:varchar(100);unique" json:"email"`
    MobileNo       string    `gorm:"type:varchar(20);unique" json:"mobile_no"`
    Password       string    `gorm:"type:text;not null" json:"password_hash"`
    Status         string    `gorm:"type:varchar(50);default:active;not null" json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (User) TableName() string {
    return "auth.users"
}

type UserActivity struct {
    ID  uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    Activity string `gorm:"type:varchar" json:"activity"`
    UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
    Count   int     `json:"count" gorm:"type:integer"`
    Month           int         `json:"month" gorm:"type:integer"`             
    Year            int         `json:"year" gorm:"type:integer"`

    User    *User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (UserActivity) TableName() string {
    return "auth.user_activities"
}

type UserResult struct {
    User *User `json:"user"`
}

type AuthUserProfile struct {
    ID                      uuid.UUID `json:"id"`
    Name                    string    `json:"name"`
    ProfilePicture          string    `json:"profile_picture"`
    Gender                  string    `json:"gender"`
    UserId                  uuid.UUID `json:"user_id"`
    CreatedAt               time.Time `json:"created_at"`
    UpdatedAt               time.Time `json:"updated_at"`
}
