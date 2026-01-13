package model

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StoreID     uuid.UUID `gorm:"type:uuid" json:"store_id"`
	Name        string    `gorm:"type:varchar" json:"name"`
	PhoneNumber string    `gorm:"type:varchar" json:"phone_number"`
	IsActive    bool      `gorm:"type:boolean" json:"is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	ModifiedAt  time.Time `gorm:"autoUpdateTime" json:"modified_at"`
}

func (Customer) TableName() string {
	return "public.customers"
}
