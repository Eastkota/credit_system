package schema

import (
	"auth_service/resolvers"
	
	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.AuthResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// "login": &graphql.Field{
			// 	Type: LoginResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"login_id": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 		"password": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return PublicAuthMiddleware(resolver.Login)(p), nil
			// 	},
			// },
			"refreshToken": &graphql.Field{
				Type: LoginResponse,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.RefreshToken(p), nil
					// return PublicAuthMiddleware(resolver.RefreshToken)(p), nil
				},
			},
			// "logout": &graphql.Field{
			// 	Type: GenericAuthResponse,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return AuthMiddleware(resolver.Logout)(p), nil
			// 	},
			// },
			"signup": &graphql.Field{
				Type: LoginResponse,
				Args: graphql.FieldConfigArgument{
					"signup_input": &graphql.ArgumentConfig{
						Type: SignupInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.Signup(p), nil
					// return PublicAuthMiddleware(resolver.Signup)(p), nil
				},
			},
			// "updatePassword": &graphql.Field{
			// 	Type: GenericAuthResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: UpdatePasswordInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return AuthMiddleware(resolver.UpdatePassword)(p), nil
			// 	},
			// },
			// "resetPassword": &graphql.Field{
			// 	Type: GenericAuthResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"user_id": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(scalar.UUID),
			// 		},
			// 		"password": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 		"confirm_password": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return PublicAuthMiddleware(resolver.ResetPassword)(p), nil
			// 	},
			// },
			// "updateSingleUserDataById": &graphql.Field{
			// 	Type: SingleUserResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"field": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 		"value": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 		"password": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return AuthMiddleware(resolver.UpdateSingleDataByID)(p), nil
			// 	},
			// },
			// "deleteUser" : &graphql.Field{
			// 	Type: SingleUserResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"userID": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(scalar.UUID),
			// 		},
			// 		"status": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 		"password": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(graphql.String),
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return AuthMiddleware(resolver.DeleteUser)(p), nil
			// 	},
			// },
			// "saveUserActivity" : &graphql.Field{
			// 	Type: UserActivityResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: UserActivityInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.SaveUserActivity(p)
			// 	},
			// },
			// "saveGameActivity" : &graphql.Field{
			// 	Type: CheyCheyActivityResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"user_id": &graphql.ArgumentConfig{
			// 			Type: graphql.NewNonNull(scalar.UUID),
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.SaveGameActivity(p)
			// 	},
			// },
		},
	})
}
