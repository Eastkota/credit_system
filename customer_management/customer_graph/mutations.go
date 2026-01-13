package graph

import (
	resolvers "credit_system/customer_management/customer_resolvers"

	"github.com/graphql-go/graphql"
)

func CustomerMutations(resolver *resolvers.CustomerResolver) graphql.Fields {
	return graphql.Fields{
		"registerCustomer": &graphql.Field{
			Type: CustomerResponse,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CustomerInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.RegisterCustomer(p), nil
			},
		},

		"addCredit": &graphql.Field{
			Type: CreditResponse,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CreditInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.AddCredit(p), nil
			},
		},
	}
}
