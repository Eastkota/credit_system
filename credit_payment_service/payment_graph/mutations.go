package graph

import (
	"credit_system/core/scalar"
	resolvers "credit_system/credit_payment_service/payment_resolvers"

	"github.com/graphql-go/graphql"
)

func NewPaymentMutationType(resolver *resolvers.PaymentResolver) graphql.Fields {
	return graphql.Fields{
		"ownerApplyPayment": &graphql.Field{
			Type: PaymentResponse,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CreatePaymentInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.PaidAmount(p), nil
			},
		},
		"ownerPaymentSubmission": &graphql.Field{
			Type: OwnerPaymentResponse,
			Args: graphql.FieldConfigArgument{
				"store_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
				"screenshot_url": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.SubmitOwnerPayment(p), nil
			},
		},
	}
}
