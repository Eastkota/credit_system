package resolvers

import (
	services "credit_system/customer_management/customer_services"
	"credit_system/customer_management/helpers"
	"credit_system/customer_management/model"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type CustomerResolver struct {
	Services services.Services
}

func NewCustomerResolver(service services.Services) *CustomerResolver {
	return &CustomerResolver{Services: service}
}

func (cr *CustomerResolver) RegisterCustomer(p graphql.ResolveParams) *model.GenericResponse {

	var customerInput model.CustomerInput
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return helpers.FormatError(fmt.Errorf("Invalid input"))
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		return helpers.FormatError(err)
	}

	err = json.Unmarshal(jsonData, &customerInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	customer, err := cr.Services.RegisterCustomer(p.Context, customerInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericResponse{
		Data: &model.CustomerResult{
			Customer: customer,
		},
		Error: nil,
	}
}

func (cr *CustomerResolver) AddCredit(p graphql.ResolveParams) *model.GenericResponse {

	var creditInput model.CreditInput
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return helpers.FormatError(fmt.Errorf("Invalid input"))
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		return helpers.FormatError(err)
	}

	err = json.Unmarshal(jsonData, &creditInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	credit, err := cr.Services.AddCredit(p.Context, creditInput)
	fmt.Printf("Credit returned from service: %+v\n", credit)
	if err != nil {
		return helpers.FormatError(err)
	}

	fmt.Printf("Credit returned from service: %+v\n", credit)

	return &model.GenericResponse{
		Data: &model.CreditResult{
			Credit: credit,
		},
		Error: nil,
	}
}

func (r *CustomerResolver) FetchCustomerBalance(p graphql.ResolveParams) *model.GenericResponse {
	customerId := p.Args["customer_id"].(uuid.UUID)
	storeId := p.Args["store_id"].(uuid.UUID)
	result, err := r.Services.FetchCustomerBalance(customerId, storeId)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericResponse{
		Data: &model.BalanceResult{
			Balance: result,
		},
		Error: nil,
	}
}

func (r *CustomerResolver) FetchAllCustomerBalance(p graphql.ResolveParams) *model.GenericResponse {
	storeId := p.Args["store_id"].(uuid.UUID)
	result, err := r.Services.FetchAllCustomerBalance(p.Context, storeId)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericResponse{
		Data: &model.MultipleBalanceResult{
			Balances: result,
		},
		Error: nil,
	}
}
