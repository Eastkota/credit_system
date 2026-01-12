package repositories

import (
    "credit_system/credit_payment_service/model"

    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
    "context"
)

type paymentRepository struct {
    DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
    return &paymentRepository{DB: db}
}

func (pr *paymentRepository) OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput ) (paymentRequest *model.CreatePaymentRequest, err error) {
    var payment model.CreatePaymentRequest

    payment.ID = uuid.New()
    payment.CustomerID = input.CustomerID
    payment.StoreID = input.StoreID
    payment.Amount = input.Amount
    payment.JournalNumber = input.JournalNumber
    payment.Status = "Pending"
    payment.CreatedAt = time.Now()
    payment.UpdatedAt = time.Now()

    result := pr.DB.WithContext(ctx).Create(&payment)
    if result.Error != nil {
        return nil, result.Error
    }

    return &payment, nil
    
}




