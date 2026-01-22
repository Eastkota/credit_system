package repositories

import (
	"context"
	"credit_system/credit_payment_service/model"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (*model.PaymentSubmission, error)
	OwnerPayment(ctx context.Context, screenshot_url, status string, storeId uuid.UUID) (*model.OwnerPaymentDetails, error)
}
