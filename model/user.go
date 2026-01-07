package model

import (
    "time"

    "github.com/google/uuid"
)


type StoreOwner struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    StoreID     uuid.UUID `gorm:"type:uuid" json:"store_id"`
    Name        string    `gorm:"type:varchar" json:"name"`
    Email       string    `gorm:"type:varchar" json:"email"`
    Password    string    `gorm:"type:varchar" json:"password_hash"`
    Status      string    `gorm:"type:varchar" json:"status"`
    PhoneNumber string    `gorm:"type:varchar" json:"phone_number"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (StoreOwner) TableName() string {
    return "public.store_owners"
}

type Store struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string    `gorm:"type:varchar" json:"name"`
    Address     string    `gorm:"type:varchar" json:"address"`
    PhoneNumber string    `gorm:"type:varchar" json:"phone_number"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Store) TableName() string {
    return "public.stores"
}


type UserResult struct {
    User *User `json:"user"`
}

