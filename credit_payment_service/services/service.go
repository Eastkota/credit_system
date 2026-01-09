package services


type Service interface {
    OwnerApplyPayment(ctx context.Context, input model.CreatePaymentInput) (model.CreatePaymentRequest, error)
}