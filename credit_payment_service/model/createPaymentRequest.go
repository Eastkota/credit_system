package model

import (
    "time"

    "github.com/google/uuid"
)


type CustomerBalances struct {
    CustomerID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"customer_id"`
    StoreID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"store_id"`

    TotalCreditGiven        string    `gorm:"type:float" json:"total_credit_given"`
    TotalPaymentsRecieved   string    `gorm:"type:float" json:"total_payments_recieved"`
    OutStandingBalance      string    `gorm:"type:float" json:"outstanding_balance"`

    LastTransactionDate     time.Time `gorm:"autoCreateTime" json:"last_transaction_date"`
    LastPaymentDate         time.Time `gorm:"autoUpdateTime" json:"last_payment_date"`
    CreatedAt               time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt               time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    
}

func (CustomerBalances) TableName() string {
    return "public.customer_balances"
}
type Credits struct {
    CustomerID                  uuid.UUID `gorm:"type:uuid;primaryKey" json:"customer_id"`
    StoreID                     uuid.UUID `gorm:"type:uuid;primaryKey" json:"store_id"`

    Amount                      *string    `gorm:"type:float" json:"amount"`

    TransactionType             *string `gorm:"type:varchar" json:"transaction_type"`
    TransactionDate             time.Time `gorm:"type:autoCreateTime" json:"transaction_date"`
    Description                 *string `gorm:"type:varchar" json:"description"`
    ItemsDescription            *string `gorm:"type:varchar" json:"items_description"`
    UnitPrices                  *string `gorm:"type:varchar" json:"unit_prices"`
    PaymentMethod               *string `gorm:"type:varchar" json:"payment_method, omitempty"`
    PaymentScreenShotURL        string `gorm:"type:varchar" json:"payment_screenshot_url"`
    ValidationStatus            string `gorm:"type:varchar" json:"validation_status"`
    ValidatedBy                 string `gorm:"varchar" json:"validated_by"`
    ValidatedAt                 time.Time `gorm:"autoCreateTime" json:"validated_at"`
    RejectionReason             string `gorm:"type:varchar" json:"rejection_reason"`
    JournalNumber               string `gorm:"type:varchar" json:"journal_number"`
    ReferenceNumber             *string `gorm:"type:varchar" json:"reference_number,omitempty"`
    CreatedBy                   *uuid.UUID `gorm:"varchar" json:"created_by,omitempty"`
    CustomerInitiated           bool `gorm:"type:boolean" json:"customer_initiated"`
    CreatedAt                   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt                   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Credits) TableName() string {
    return "public.credits"
}

type CreatePaymentRequest struct {
    ID                          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    CustomerID                  uuid.UUID `gorm:"type:uuid" json:"customer_id"`
    StoreID                     uuid.UUID `gorm:"type:uuid;primaryKey" json:"store_id"`

    Amount                      string    `gorm:"type:float" json:"amount"`
    JournalNumber               string    `gorm:"type:varchar" json:"journal_number"`
    Status                     string    `gorm:"type:varchar" json:"status"`

    CreatedAt                   time.Time `json:"created_at,omitempty"`
    UpdatedAt                   time.Time `json:"updated_at,omitempty"` 
}
func (CreatePaymentRequest) TableName() string {
    return "public.payment_submissions"
}






