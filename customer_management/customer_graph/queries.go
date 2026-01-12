package graph

import (
	"credit_system/core/schema"
	resolvers "credit_system/customer_management/customer_resolvers"
	"credit_system/customer_management/helpers"
	"credit_system/customer_management/model"

	"github.com/graphql-go/graphql"
)

func CustomerQueries(resolver *resolvers.CustomerResolver) graphql.Fields {
	return graphql.Fields{
		"customerService": &graphql.Field{
			Type: graphql.NewNonNull(schema.ServiceInfoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				schema, err := schema.GetSchema()
				if err != nil {
					return nil, err
				}

				serviceInfo := model.Service{
					Name:    "CustomerService",
					Version: "1.0.0",
					Schema:  helpers.ConvertSchemaToString(schema),
				}
				return serviceInfo, nil
			},
		},
	}
}
