package graph

import (
	"credit_system/core/scalar"
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
		"FetchCustomerBalance": &graphql.Field{
			Type: BalanceResponse,
			Args: graphql.FieldConfigArgument{
				"customer_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
				"store_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchCustomerBalance(p), nil
			},
		},
		"FetchAllCustomerBalance": &graphql.Field{
			Type: MultipleBalanceResponse,
			Args: graphql.FieldConfigArgument{
				"store_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchAllCustomerBalance(p), nil
			},
		},
		"FetchAllCustomerByStoreId": &graphql.Field{
			Type: MultipleCustomerResponse,
			Args: graphql.FieldConfigArgument{
				"store_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.AuthMiddleware[model.GenericCustomerResponse](resolver.FetchAllCustomersByStoreId)(p), nil
			},
		},

		"FetchCustomerById": &graphql.Field{
			Type: CustomerResponse,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(scalar.UUID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchCustomerById(p), nil
			},
		},
	}
}
