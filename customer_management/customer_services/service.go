package services

import (
	"credit_system/customer_management/model"

	"context"

	"github.com/google/uuid"
)

type Services interface {
	RegisterCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error)
	AddCredit(ctx context.Context, input model.CreditInput) (*model.Credit, error)
	FetchCustomerBalance(customer_id, store_id uuid.UUID) (*model.BalanceUpdate, error)
	FetchAllCustomerBalance(ctx context.Context, store_id uuid.UUID) ([]model.BalanceUpdate, error)
	FetchAllCustomerByStoreId(ctx context.Context, store_id uuid.UUID) ([]model.Customer, error)
	FetchCustomerById(ctx context.Context, customer_id uuid.UUID) (*model.Customer, error)
}
