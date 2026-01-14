package helpers

import "credit_system/customer_management/model"

func FormatError(err error) *model.GenericResponse {
	return &model.GenericResponse{
		Data: nil,
		Error: &model.Error{
			Message: err.Error(),
		},
	}
}
