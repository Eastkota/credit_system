package resolvers

import (
	"credit_system/customer_management/customer_services"
	"credit_system/customer_management/model"

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
	input := model.CustomerInput{
		StoreID:     p.Args["input"].(map[string]interface{})["store_id"].(uuid.UUID),
		Name:        p.Args["input"].(map[string]interface{})["name"].(string),
		PhoneNumber: p.Args["input"].(map[string]interface{})["phone_number"].(string),
		Email:       p.Args["input"].(map[string]interface{})["email"].(string),
		CreditLimit: p.Args["input"].(map[string]interface{})["credit_limit"].(float64),
		IsActive:    true,
	}

	customer, err := cr.Services.RegisterCustomer(p.Context, input)
	if err != nil {
		return &model.GenericResponse{
			Data: nil,
			Error: &model.Error{
				Message: err.Error(),
			},
		}
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
