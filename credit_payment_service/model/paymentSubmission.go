package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentSubmission struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"customer_id"`
	StoreID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"store_id"`
	ValidatedByOwnerId uuid.UUID `gorm:"type:uuid" json:"validated_by_owner_id"`
	LinkedCreditID     uuid.UUID `gorm:"type:uuid" json:"linked_credit_id"`

	Amount      float64   `gorm:"type:float" json:"amount"`
	Status      string    `gorm:"type:varchar" json:"status"`
	SubmittedAt time.Time `gorm:"autoCreateTime" json:"submitted_at"`
	ValidatedAt time.Time `gorm:"autoUpdateTime" json:"validated_at"`

	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	ModifiedAt time.Time `gorm:"autoUpdateTime" json:"modified_at"`
}

func (PaymentSubmission) TableName() string {
	return "public.payment_submissions"
}

type CustomerBalance struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"customer_id"`
	TotalPayementsReceived float64   `gorm:"type:float" json:"total_payements_received"`
	LastPaymentDate        time.Time `gorm:"autoCreateTime" json:"last_payment_date"`
}

func (CustomerBalance) TableName() string {
	return "public.customer_balances"
}
