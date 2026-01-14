package model

type GenericPaymentResponse struct {
	Data  interface{}   `json:"data,omitempty"`
	Error *PaymentError `json:"error,omitempty"`
}

type PaymentSubmissionResponse struct {
	PaymentSubmission PaymentSubmission `json:"create_payment_request"`
}
