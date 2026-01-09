package model


type CreatePaymentInput struct {
    StoreID    string `json:"store_id"`
    CustomerID string `json:"customer_id"`
    Amount     string `json:"amount"`
    JournalNumber string `json:"journal_number"`
    Status     string `json:"status"`
}


