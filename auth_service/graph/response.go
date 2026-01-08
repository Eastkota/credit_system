package schema

import "github.com/graphql-go/graphql"

var LoginResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: Login},
		"error": &graphql.Field{Type: AuthError},
	},
})

var GenericAuthResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAuthResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: AuthGenericSuccessData},
		"error": &graphql.Field{Type: AuthError},
	},
})

var CheckForExistingUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckForExistingUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: ExistUser},
		"error": &graphql.Field{Type: AuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var StoreResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "StoreResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: StoreResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var ValidateTokenResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "ValidateTokenResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: ValidateTokenResult},
		"error": &graphql.Field{Type: AuthError},
	},
})




