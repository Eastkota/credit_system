package resolvers

import (
	"credit_system/credit_payment_service/helpers"
	"credit_system/credit_payment_service/model"
	services "credit_system/credit_payment_service/payment_services"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type PaymentResolver struct {
	Services services.Services // Inject Services
}

func NewPaymentResolver(service services.Services) *PaymentResolver {
	return &PaymentResolver{Services: service}
}

func (pr *PaymentResolver) PaidAmount(p graphql.ResolveParams) *model.GenericPaymentResponse {
	var paymentInput model.CreatePaymentInput

	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return helpers.FormatError(fmt.Errorf("Invalid Input"))
	}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return helpers.FormatError(fmt.Errorf("Failed to parse input: %w", err))
	}
	err = json.Unmarshal(jsonData, &paymentInput)
	if err != nil {
		return helpers.FormatError(fmt.Errorf("Failed to parse input data: %w", err))
	}

	paymentSubmission, err := pr.Services.OwnerApplyPayment(p.Context, paymentInput)
	fmt.Println("Payment Submission:", paymentSubmission)
	if err != nil {
		return helpers.FormatError(fmt.Errorf("Failed to apply payment: %w", err))
	}

	return &model.GenericPaymentResponse{
		Data: &model.PaymentSubmissionResponse{
			PaymentSubmission: *paymentSubmission,
		},
		Error: nil,
	}
}

func (r *PaymentResolver) SubmitOwnerPayment(p graphql.ResolveParams) *model.GenericPaymentResponse {
	storeId := p.Args["store_id"].(uuid.UUID)
	screenshot_url := p.Args["screenshot_url"].(string)
	status := p.Args["status"].(string)
	result, err := r.Services.OwnerPayment(p.Context, screenshot_url, status, storeId)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericPaymentResponse{
		Data: &model.OwnerPaymentSubmissionResponse{
			OwnerPaymentSubmission: *result,
		},
		Error: nil,
	}
}
