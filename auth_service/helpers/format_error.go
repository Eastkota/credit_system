package helpers

import "auth_service/model"

func FormatError(err error) *model.GenericAuthResponse {
	return &model.GenericAuthResponse{
		Data: nil,
		Error: &model.AuthError{
			Message: err.Error(),
		},
	}
}
