package schema

import (
	"auth_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: scalar.UUID},
		"user_identifier": &graphql.Field{Type: graphql.String},
		"email":           &graphql.Field{Type: graphql.String},
		"mobile_no":       &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"created_at":      &graphql.Field{Type: scalar.Time},
		"updated_at":      &graphql.Field{Type: scalar.Time},
	},
})

// Define UserType
var AuthUserProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUserProfile",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})

// Define AccessToken
var AccessToken = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccessToken",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: scalar.UUID},
		"identifier": &graphql.Field{Type: graphql.String},
		"user_id":    &graphql.Field{Type: scalar.UUID},
		"client_id":  &graphql.Field{Type: graphql.String},
		"scopes":     &graphql.Field{Type: graphql.NewList(graphql.String)},
		"expires_at": &graphql.Field{Type: scalar.Time},
	},
})

// Define Token
var Token = graphql.NewObject(graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"token_type":    &graphql.Field{Type: graphql.String},
		"access_token":  &graphql.Field{Type: graphql.String},
		"refresh_token": &graphql.Field{Type: graphql.String},
		"expires_in":    &graphql.Field{Type: scalar.Time},
	},
})

var Login = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: User},
		"profile":    &graphql.Field{Type: AuthUserProfile},
		"token":      &graphql.Field{Type: Token},
		"membership": &graphql.Field{Type: AuthUserMembership},
	},
})

var ExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})
var UserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: User},
	},
})

var AuthGenericSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthGenericSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var ValidateTokenResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "ValidateTokenResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: User},
	},
})

var UserActivity = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserActivity",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: scalar.UUID},
		"user_id":    &graphql.Field{Type: scalar.UUID},
		"activity":  &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: scalar.Time},
		"updated_at": &graphql.Field{Type: scalar.Time},
		"count":      &graphql.Field{Type: graphql.Int},
		"month":      &graphql.Field{Type: graphql.Int},
		"year":       &graphql.Field{Type: graphql.Int},

		"user":		&graphql.Field{Type: User},
	},
})

var UserActivityResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserActivityResult",
	Fields: graphql.Fields{
		"user_activity": &graphql.Field{Type: UserActivity},
	},
})
