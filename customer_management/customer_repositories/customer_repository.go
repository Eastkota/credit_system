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
		ID:          uuid.New(),
		StoreID:     input.StoreID,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		IsActive:    true,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	if err := repo.DB.WithContext(ctx).Create(customer).Error; err != nil {
		return nil, fmt.Errorf("failed to create customer: %v", err)
	}

	return customer, nil
}

func (repo *CustomerRepository) AddCredit(ctx context.Context, input model.CreditInput) (*model.Credit, error) {
	var transactionType = input.TransactionType
	var credit_amount float64

	// Create credit record
	credit := &model.Credit{
		ID:               uuid.New(),
		StoreId:          input.StoreId,
		CustomerId:       input.CustomerId,
		Amount:           input.Amount,
		TransactionType:  transactionType,
		TransactionDate:  time.Now(),
		ItemsDescription: input.ItemsDescription,
		JournalNumber:    input.JournalNumber,
		CreatedAt:        time.Now(),
		ModifiedAt:       time.Now(),
	}

	if err := repo.DB.WithContext(ctx).Create(credit).Error; err != nil {
		return nil, fmt.Errorf("failed to create credit: %v", err)
	}

	if transactionType == "credit_given" {
		credit_amount = input.Amount
	} else {
		credit_amount = -input.Amount
	}

	existingBalance, err := repo.FetchCustomerBalance(input.CustomerId, input.StoreId)

	if err != nil {
		newBalance := &model.BalanceUpdate{
			Id:                    uuid.New(),
			CustomerId:            input.CustomerId,
			StoreId:               input.StoreId,
			OutstandingBalance:    credit_amount,
			TotalCreditGiven:      credit_amount,
			TotalPaymentsReceived: 0.00,
			LastPaymentDate:       time.Now(),
			LastCreditDate:        time.Now(),
			LastTransactionDate:   time.Now(),
			CreatedAt:             time.Now(),
			ModifiedAt:            time.Now(),
		}

		if err := repo.DB.WithContext(ctx).Create(newBalance).Error; err != nil {
			return nil, fmt.Errorf("failed to create balance: %v", err)
		}
	} else {
		updates := map[string]interface{}{
			"outstanding_balance":   existingBalance.OutstandingBalance + credit_amount,
			"last_transaction_date": time.Now(),
			"modified_at":           time.Now(),
		}

		if transactionType == "credit_given" {
			updates["last_credit_date"] = time.Now()
			updates["total_credit_given"] = existingBalance.TotalCreditGiven + credit_amount
		} else {
			updates["last_payment_date"] = time.Now()
			updates["total_payments_received"] = existingBalance.TotalPaymentsReceived + input.Amount
		}

		if err := repo.DB.WithContext(ctx).Model(&model.BalanceUpdate{}).
			Where("customer_id = ? AND store_id = ?", input.CustomerId, input.StoreId).
			Updates(updates).Error; err != nil {
			return nil, fmt.Errorf("failed to update balance: %v", err)
		}
	}

	return credit, nil
}

func (repo *CustomerRepository) FetchCustomerBalance(customer_id, store_id uuid.UUID) (*model.BalanceUpdate, error) {
	var customer_balances model.BalanceUpdate
	result := repo.DB.Where("customer_id = ? AND store_id = ?", customer_id, store_id).
		Find(&customer_balances)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch customer balance: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("customer balance record not found")
	}
	return &customer_balances, nil
}

func (repo *CustomerRepository) FetchAllCustomerBalance(ctx context.Context, store_id uuid.UUID) ([]model.BalanceUpdate, error) {
	var results []model.BalanceUpdate
	result := repo.DB.WithContext(ctx).
		Where("store_id = ?", store_id).
		Order("outstanding_balance DESC").
		Find(&results)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get customer balances: %v", result.Error)
	}

	return results, nil
}

func (repo *CustomerRepository) FetchAllCustomerByStoreId(ctx context.Context, store_id uuid.UUID) ([]model.Customer, error) {
	var results []model.Customer
	result := repo.DB.WithContext(ctx).
		Where("store_id = ?", store_id).
		Find(&results)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get customer balances: %v", result.Error)
	}

	return results, nil
}
func (repo *CustomerRepository) FetchCustomerById(ctx context.Context, customer_id uuid.UUID) (*model.Customer, error) {
	var customer model.Customer
	result := repo.DB.WithContext(ctx).
		Where("id = ?", customer_id).
		Find(&customer)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get customer balances: %v", result.Error)
	}

	return &customer, nil
}
