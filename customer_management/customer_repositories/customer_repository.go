package repositories

import (
	"context"
	"credit_system/customer_management/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (repo *CustomerRepository) RegisterCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	customer := &model.Customer{
		ID:           uuid.New(),
		StoreID:      input.StoreID,
		Name:         input.Name,
		PhoneNumber:  input.PhoneNumber,
		Email:        input.Email,
		Credit_limit: input.CreditLimit,
		HasCredit:    true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := repo.DB.WithContext(ctx).Create(customer).Error; err != nil {
		return nil, fmt.Errorf("failed to create customer: %v", err)
	}

	return customer, nil
}
