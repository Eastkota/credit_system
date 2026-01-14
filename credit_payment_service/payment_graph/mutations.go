package graph

import (
	"credit_system/credit_payment_service/payment_resolvers"
	
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

	}
}

