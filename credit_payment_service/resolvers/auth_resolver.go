package resolvers

import (
	"credit_system/credit_payment_service/helpers"
	"credit_system/credit_payment_service/model"
	"credit_system/credit_payment_service/services"

	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type PaymentResolver struct {
	Services services.Services // Inject Services
}

func NewPaymentResolver(service services.Services) *PaymentResolver {
	return &PaymentResolver{Services: service}
}

func (pr *PaymentResolver) PaidAmount(p graphql.ResolveParams) *model.GenericPaymentResponse {
	// If you use input object:
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return helpers.FormatPaymentError(errors.New("missing input"))
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return helpers.FormatError(err)
	}

	var paidAmountInput model.CreatePaymentInput
	err = json.Unmarshal(inputBytes, &paidAmountInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	paymentRequest, err := pr.Services.PaymentService.OwnerApplyPayment(p.Context, paidAmountInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericPaymentResponse{
		Data: &model.PaymentSuccessData{
			Payment: paymentRequest,
		},
		Error: nil,
	}

}
