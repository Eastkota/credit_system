package model

import (
	"time"

	"github.com/google/uuid"
)

type CustomerInput struct {
	StoreID     uuid.UUID `json:"store_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
}

type CreditInput struct {
	StoreId          uuid.UUID `json:"store_id"`
	CustomerId       uuid.UUID `json:"customer_id"`
	Amount           float64   `json:"amount"`
	TransactionType  string    `json:"transaction_type"`
	TransactionDate  time.Time `json:"transaction_date"`
	ItemsDescription string    `json:"items_description"`
	JournalNumber    string    `json:"journal_number"`
}
