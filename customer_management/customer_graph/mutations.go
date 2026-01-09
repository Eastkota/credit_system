package graph

import (
	resolvers "credit_system/customer_management/customer_resolvers"

	"github.com/graphql-go/graphql"
)

func NewCustomerMutationType(resolver *resolvers.CustomerResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
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
		},
	})
}
