package graph

import (
	"credit_system/core/schema"
	resolvers "credit_system/customer_management/customer_resolvers"
	"credit_system/customer_management/model"

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
				middleware := schema.AuthMiddleware[model.GenericCustomerResponse](resolver.RegisterCustomer)
				return middleware(p), nil
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
				middleware := schema.AuthMiddleware[model.GenericCustomerResponse](resolver.AddCredit)
				return middleware(p), nil
			},
		},
	}
}
