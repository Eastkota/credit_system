package generic_response

import (
    "credit_system/auth_service/model"
	payment_model "credit_system/credit_payment_service/model"
	customer_model "credit_system/customer_management/model"
)

// FormatMiddlewareError detects the type T and returns the appropriate pointer 
// with the error field populated.
func FormatMiddlewareError[T any](err error) *T {
    res := new(T)
    errMessage := err.Error()

    // Use a type switch on the pointer to identify which struct we are filling
    switch any(res).(type) {
    case *model.GenericAuthResponse:
        any(res).(*model.GenericAuthResponse).Error = &model.AuthError{Message: errMessage}
    case *payment_model.GenericPaymentResponse:
        any(res).(*payment_model.GenericPaymentResponse).Error = &payment_model.PaymentError{Message: errMessage}
    case *customer_model.GenericCustomerResponse:
        any(res).(*customer_model.GenericCustomerResponse).Error = &customer_model.Error{Message: errMessage}
    }

    return res
}