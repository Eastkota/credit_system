package repositories

import (
	"credit_system/credit_payment_service/model"
    "context"
)

type PaymentRepository interface {
    OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput ) (paymentRequest *model.CreatePaymentRequest, err error)
}