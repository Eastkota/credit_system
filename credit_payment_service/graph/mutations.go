package schema

import (
	"credit_system/auth_service/resolvers"
	
	"github.com/graphql-go/graphql"
)

func PaymentMutations(resolver *resolvers.AuthResolver) graphql.Fields {
	return graphql.Fields{
		"ownerApplyPayment": &graphql.Field{
			Type: PaymentResponse,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CreatePaymentInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.OwnerApplyPayment(p)
			}
		},

	},
}

