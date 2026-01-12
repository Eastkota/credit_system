package model

import (
    "github.com/google/uuid"
)

type CreatePaymentInput struct {
    StoreID    uuid.UUID `json:"store_id"`
    CustomerID uuid.UUID `json:"customer_id"`
    Amount     string `json:"amount"`
    JournalNumber string `json:"journal_number"`
    Status     string `json:"status"`
}


