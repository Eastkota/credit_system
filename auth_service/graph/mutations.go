package graph

import (
	"credit_system/auth_service/resolvers"
	
	"github.com/graphql-go/graphql"
)

func AuthMutations(resolver *resolvers.AuthResolver) graphql.Fields {
    return graphql.Fields{
        "login": &graphql.Field{
            Type: LoginResponse,
            Args: graphql.FieldConfigArgument{
                "phone_number": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "password": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return PublicAuthMiddleware(resolver.Login)(p), nil
            },
        },

        "refreshToken": &graphql.Field{
            Type: LoginResponse,
            Args: graphql.FieldConfigArgument{
                "token": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return resolver.RefreshToken(p), nil
            },
        },

        "signup": &graphql.Field{
            Type: LoginResponse,
            Args: graphql.FieldConfigArgument{
                "signup_input": &graphql.ArgumentConfig{
                    Type: SignupInput,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return resolver.Signup(p), nil
            },
        },

        "updateSingleUserDataById": &graphql.Field{
            Type: SingleUserResponse,
            Args: graphql.FieldConfigArgument{
                "field": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "value": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "password": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return AuthMiddleware(resolver.UpdateSingleDataByID)(p), nil
            },
        },
    }
}
