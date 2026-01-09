package model

import "github.com/google/uuid"

type CustomerInput struct {
	StoreID     uuid.UUID `json:"store_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreditLimit float64   `json:"credit_limit"`
	IsActive    bool      `json:"is_active"`
}
