package schema

import (
	"credit_system/auth_service/helpers"
	"credit_system/auth_service/model"
	"credit_system/auth_service/resolvers"
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

func NewQueryType(resolver *resolvers.AuthResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
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
			"FetchOwnerByID": &graphql.Field{
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
		},
	})
}
