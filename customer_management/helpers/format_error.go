package helpers

import "credit_system/customer_management/model"

func FormatError(err error) *model.GenericCustomerResponse {
	return &model.GenericCustomerResponse{
		Data: nil,
		Error: &model.Error{
			Message: err.Error(),
		},
	}
}
