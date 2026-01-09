package services

import (
	"context"
	"credit_system/customer_management/customer_repositories"
	"credit_system/customer_management/model"
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
