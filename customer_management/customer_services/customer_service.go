package services

import (
	"context"
	repositories "credit_system/customer_management/customer_repositories"
	"credit_system/customer_management/model"

	"github.com/google/uuid"
)

type CustomerService struct {
	Repository repositories.Repository
}

func NewCustomerService(repository repositories.Repository) *CustomerService {
	return &CustomerService{Repository: repository}
}

func (cs *CustomerService) RegisterCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	return cs.Repository.RegisterCustomer(ctx, input)
}

func (cs *CustomerService) AddCredit(ctx context.Context, input model.CreditInput) (*model.Credit, error) {
	return cs.Repository.AddCredit(ctx, input)
}
func (cs *CustomerService) FetchCustomerBalance(customer_id, store_id uuid.UUID) (*model.BalanceUpdate, error) {
	return cs.Repository.FetchCustomerBalance(customer_id, store_id)
}
func (cs *CustomerService) FetchAllCustomerBalance(ctx context.Context, store_id uuid.UUID) ([]model.BalanceUpdate, error) {
	return cs.Repository.FetchAllCustomerBalance(ctx, store_id)
}
func (cs *CustomerService) FetchAllCustomerByStoreId(ctx context.Context, store_id uuid.UUID) ([]model.Customer, error) {
	return cs.Repository.FetchAllCustomerByStoreId(ctx, store_id)
}
func (cs *CustomerService) FetchCustomerById(ctx context.Context, customer_id uuid.UUID) (*model.Customer, error) {
	return cs.Repository.FetchCustomerById(ctx, customer_id)
}
