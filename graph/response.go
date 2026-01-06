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

var ValidateTokenResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "ValidateTokenResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: ValidateTokenResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var UserActivityResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserActivityResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserActivityResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var CheyCheyActivityResult = graphql.NewObject(graphql.ObjectConfig{
    Name: "CheyCheyActivityResult",
    Fields: graphql.Fields{
        "message": &graphql.Field{
            Type: graphql.String,
        },
    },
})

var CheyCheyActivityResponse = graphql.NewObject(graphql.ObjectConfig{
    Name: "CheyCheyActivityResponse",
    Fields: graphql.Fields{
        "data":  &graphql.Field{Type: CheyCheyActivityResult},
        "error": &graphql.Field{Type: AuthError},
    },
})

