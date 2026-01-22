package services

import (
	"credit_system/credit_payment_service/model"
	repositories "credit_system/credit_payment_service/payment_repositories"

	"context"

	"github.com/google/uuid"
)

type PaymentService struct {
	Repository repositories.PaymentRepository
}

func NewPaymentService(repository repositories.PaymentRepository) *PaymentService {
	return &PaymentService{Repository: repository}
}

func (as *PaymentService) OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (*model.PaymentSubmission, error) {
	return as.Repository.OwnerApplyPayment(ctx, input)
}

func (as *PaymentService) OwnerPayment(ctx context.Context, screenshot_url, status string, storeId uuid.UUID) (*model.OwnerPaymentDetails, error) {
	return as.Repository.OwnerPayment(ctx, screenshot_url, status, storeId)
}
