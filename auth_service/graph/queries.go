package graph


import (
	"credit_system/auth_service/helpers"
	"credit_system/auth_service/model"
	"credit_system/auth_service/resolvers"
	"credit_system/core/schema"
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)


func AuthQueries(resolver *resolvers.AuthResolver) graphql.Fields {
	return graphql.Fields{
		"service": &graphql.Field{
			Type: graphql.NewNonNull(Service),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				s, err := schema.GetSchema()
				if err != nil {
					return nil, err
				}

				return model.Service{
					Name:    "AuthService",
					Version: "1.0.0",
					Schema:  helpers.ConvertSchemaToString(s),
				}, nil
			},
		},

		"checkForExistingUser": &graphql.Field{
			Type: CheckForExistingUserResponse,
			Args: graphql.FieldConfigArgument{
				"field": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"value": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.CheckForExistingUser(p), nil
			},
		},

		"fetchOwnerByID": &graphql.Field{
			Type: SingleUserResponse,
			Args: graphql.FieldConfigArgument{
				"owner_id": &graphql.ArgumentConfig{
					Type: scalar.UUID,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchOwnerByID(p), nil
			},
		},

		"fetchStore": &graphql.Field{
			Type: StoreResponse,
			Args: graphql.FieldConfigArgument{
				"store_id": &graphql.ArgumentConfig{
					Type: scalar.UUID,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchStore(p), nil
			},
		},

		"fetchStoreByOwnerID": &graphql.Field{
			Type: StoreResponse,
			Args: graphql.FieldConfigArgument{
				"owner_id": &graphql.ArgumentConfig{
					Type: scalar.UUID,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.FetchStoreByOwnerID(p), nil
			},
		},

		"validateToken": &graphql.Field{
			Type: ValidateTokenResponse,
			Args: graphql.FieldConfigArgument{
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.ValidateToken(p), nil
			},
		},
	}
}
