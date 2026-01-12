package resolvers

import (
	"credit_system/customer_management/customer_services"
	"credit_system/customer_management/helpers"
	"credit_system/customer_management/model"
	"encoding/json"

	// "github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type CustomerResolver struct {
	Services services.Services
}

func NewCustomerResolver(service services.Services) *CustomerResolver {
	return &CustomerResolver{Services: service}
}

func (cr *CustomerResolver) RegisterCustomer(p graphql.ResolveParams) *model.GenericResponse {

	// input := model.CustomerInput{
	// 	StoreID:     p.Args["store_id"].(uuid.UUID),
	// 	Name:        p.Args["name"].(string),
	// 	PhoneNumber: p.Args["phone_number"].(string),
	// 	HasCredit:   p.Args["has_credit"].(bool),
	// }
	var customerInput model.CustomerInput
	input := p.Args["input"].(map[string]interface{})

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

// func (cr *CustomerResolver) FetchStore(p graphql.ResolveParams) *model.GenericResponse {
// 	storeID := p.Args["store_id"].(uuid.UUID)
// 	result, err := cr.Services.FetchStore(storeID)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.StoreResult{
// 			Store: result,
// 		},
// 		Error: nil,
// 	}
// }
