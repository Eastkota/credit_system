package graph

import (
	"credit_system/core/schema"
	resolvers "credit_system/customer_management/customer_resolvers"
	"credit_system/customer_management/helpers"
	"credit_system/customer_management/model"

	"github.com/graphql-go/graphql"
)

func NewCustomerQueryType(resolver *resolvers.CustomerResolver) graphql.Fields {
	return graphql.Fields{
		"service": &graphql.Field{
			Type: graphql.NewNonNull(Service),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				schema, err := schema.GetSchema()
				if err != nil {
					return nil, err
				}

				serviceInfo := model.Service{
					Name:    "AuthService",
					Version: "1.0.0",
					Schema:  helpers.ConvertSchemaToString(schema),
				}
				return serviceInfo, nil
			},
		},
	}
}
