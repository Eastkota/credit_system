package repositories

import (
	"credit_system/credit_payment_service/model"
	"fmt"

	"time"

	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentSubmissionRepository struct {
	DB *gorm.DB
}

func NewPaymentSubmissionRepository(db *gorm.DB) *PaymentSubmissionRepository {
	return &PaymentSubmissionRepository{DB: db}
}

func (pr *PaymentSubmissionRepository) OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (*model.PaymentSubmission, error) {

	payment := model.PaymentSubmission{
		ID:                 uuid.New(),
		CustomerID:         input.CustomerID,
		StoreID:            input.StoreID,
		ValidatedByOwnerId: input.ValidatedByOwnerId,
		LinkedCreditID:     input.LinkedCreditID,

		Amount:      input.Amount,
		Status:      input.Status,
		ValidatedAt: time.Now(),

		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}
	if err := pr.DB.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, fmt.Errorf("failed to create payment submission: %w", err)
	}
	return &payment, nil
}
