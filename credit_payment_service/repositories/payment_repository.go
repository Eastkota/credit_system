package repositories

import (
    "credit_system/credit_payment_service/model"

    "fmt"
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
    "context"
)

type PaymentRepository struct {
    DB *gorm.DB
}

func NewRepository(db *gorm.DB) *PaymentRepository {
    return &PaymentRepository{DB: db}
}

func (pr *PaymentRepository) OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput ) (paymentRequest *model.CreatePaymentRequest, err error) {
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




