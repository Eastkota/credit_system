package services

import (
	"context"
	"credit_system/customer_management/model"
)

type Services interface {
	RegisterCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error)
}
