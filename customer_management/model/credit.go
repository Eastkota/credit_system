package model

import (
	"time"

	"github.com/google/uuid"
)

type Credit struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StoreId          uuid.UUID `gorm:"type:uuid; not null" json:"store_id"`
	CustomerId       uuid.UUID `gorm:"type:uuid; not null" json:"customer_id"`
	Amount           float64   `gorm:"type:decimal" json:"amount"`
	TransactionType  string    `gorm:"type:varchar(255)" json:"transaction_type"`
	TransactionDate  time.Time `grom:"type:timestamptz" json:"transaction_date"`
	ItemsDescription string    `grom:"type:varchar(500)" json:"items_description"`
	JournalNumber    string    `grom:"type:varchar(255)" json:"journal_number"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	ModifiedAt       time.Time `gorm:"autoUpdateTime" json:"modified_at"`
}

func (Credit) TableName() string {
	return "public.credits"
}

type BalanceUpdate struct {
	Id                    uuid.UUID `grom:"type:uuid;primaryKey" json:"id"`
	CustomerId            uuid.UUID `grom:"type:uuid; not null" json:"customer_id"`
	StoreId               uuid.UUID `grom:"type:uuid; not null" json:"store_id"`
	TotalCreditGiven      float64   `grom:"type:decimal" json:"total_credit_given"`
	TotalPaymentsReceived float64   `grom:"type:decimal" json:"total_payments_received"`
	OutstandingBalance    float64   `grom:"type:decimal" json:"outstanding_balance"`
	LastCreditDate        time.Time `grom:"type:timestamptz" json:"last_credit_date"`
	LastPaymentDate       time.Time `grom:"type:timestamptz" json:"last_payment_date"`
	LastTransactionDate   time.Time `grom:"type:timestamptz" json:"last_transaction_date"`
	CreatedAt             time.Time `gorm:"type:timestamptz" json:"created_at"`
	ModifiedAt            time.Time `gorm:"type:timestamptz" json:"modified_at"`
}

func (BalanceUpdate) TableName() string {
	return "public.customer_balances"
}
