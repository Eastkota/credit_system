package services

import (
	"credit_system/config"
	"credit_system/credit_payment_service/model"
	"credit_system/credit_payment_service/repositories"

	"fmt"
	"time"
	"strings"
    "context"


	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
    ErrInvalidInput     = fmt.Errorf("invalid_input")
    ErrOverPayment      = fmt.Errorf("over_payment")
    ErrBalanceNotFound  = fmt.Errorf("balance_not_found")
)

type PaymentService struct {
	Repository repositories.PaymentRepository 
}

func NewPaymentService(repository repositories.PaymentRepository) *PaymentService {
	return &PaymentService{Repository: repository}
}

func (as *PaymentService) OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput)(model.CreatePaymentRequest, error) {
    paymentRequest, err := as.Repository.OwnerApplyPayment(ctx, input)
    if err != nil {
        return model.CreatePaymentRequest{}, err
    }

    return *paymentRequest, nil
}