package repositories

import (
	"context"
	"credit_system/customer_management/model"
)

type Repository interface {
	RegisterCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error)
}
