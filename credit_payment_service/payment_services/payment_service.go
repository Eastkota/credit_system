package services

import (
	"credit_system/credit_payment_service/model"
	repositories "credit_system/credit_payment_service/payment_repositories"

	"context"
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
