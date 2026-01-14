package repositories

import (
	"context"
	"credit_system/credit_payment_service/model"
)

type PaymentRepository interface {
	OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (*model.PaymentSubmission, error)
}
