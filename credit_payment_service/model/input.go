package model

import (
	"time"

	"github.com/google/uuid"
)

type CreatePaymentInput struct {
	ID                 uuid.UUID `json:"id"`
	CustomerID         uuid.UUID `json:"customer_id"`
	StoreID            uuid.UUID `json:"store_id"`
	ValidatedByOwnerId uuid.UUID `json:"validated_by_owner_id"`
	LinkedCreditID     uuid.UUID `json:"linked_credit_id"`

	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	SubmittedAt time.Time `json:"submitted_at"`
	ValidatedAt time.Time `json:"validated_at"`

	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
