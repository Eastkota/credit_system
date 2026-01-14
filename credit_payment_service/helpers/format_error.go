package helpers

import "credit_system/credit_payment_service/model"

func FormatError(err error) *model.GenericPaymentResponse {
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.PaymentError{
			Message: err.Error(),
		},
	}
}
