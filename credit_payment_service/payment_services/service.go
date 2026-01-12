package services

import (
    "context"
    "credit_system/credit_payment_service/model"
)

type Services interface {
    OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (model.CreatePaymentRequest, error)
}