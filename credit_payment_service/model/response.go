package model

type GenericPaymentResponse struct {
    Data interface{} `json:"data,omitempty"`
    Error *PaymentError   `json:"error,omitempty"`
}

type PaymentSuccessData struct {
    Payment CreatePaymentRequest `json:"create_payment_request"`
}