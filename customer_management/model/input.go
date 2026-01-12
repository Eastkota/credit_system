package model

import "github.com/google/uuid"

type CustomerInput struct {
	StoreID     uuid.UUID `json:"store_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	HasCredit   bool      `json:"has_credit"`
}
