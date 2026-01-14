package model

import (
    "time"

    "github.com/google/uuid"
)


type StoreOwner struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string    `gorm:"type:varchar" json:"name"`
    Password    string    `gorm:"type:varchar" json:"password"`
    Status      string    `gorm:"type:varchar" json:"status"`
    PhoneNumber string    `gorm:"type:varchar" json:"phone_number"`
    AccountNumber string    `gorm:"type:varchar" json:"account_number"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
 
func (StoreOwner) TableName() string {
    return "public.store_owners"
}

type Store struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string    `gorm:"type:varchar" json:"name"`
    OwnerID     uuid.UUID `gorm:"type:uuid" json:"owner_id"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    Owner       *StoreOwner `gorm:"foreignKey:OwnerID" json:"owner"`
}

func (Store) TableName() string {
    return "public.stores"
}


