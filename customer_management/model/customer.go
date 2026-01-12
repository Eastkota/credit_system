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
	HasCredit   bool      `gorm:"type:boolean" json:"has_credit"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Customer) TableName() string {
	return "public.customers"
}
