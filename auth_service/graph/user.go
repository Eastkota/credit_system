package schema

import (
	"auth_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var StoreOwner = graphql.NewObject(graphql.ObjectConfig{
	Name: "StoreOwner",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"phone_number":               &graphql.Field{Type: graphql.String},
		"password":                   &graphql.Field{Type: graphql.String},
		"status":                     &graphql.Field{Type: graphql.String},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})

// Define UserType
var Store = graphql.NewObject(graphql.ObjectConfig{
	Name: "Store",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: scalar.UUID},
		"name": &graphql.Field{Type: graphql.String},
		"owner_id": &graphql.Field{Type: scalar.UUID},
		"created_at": &graphql.Field{Type: scalar.Time},
		"updated_at": &graphql.Field{Type: scalar.Time},
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
		"owner":       &graphql.Field{Type: StoreOwner},
		"store":    &graphql.Field{Type: Store},
		"token":      &graphql.Field{Type: Token},
	},
})

var ExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"owner_id":    &graphql.Field{Type: scalar.UUID},
	},
})
var UserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResult",
	Fields: graphql.Fields{
		"owner": &graphql.Field{Type: StoreOwner},
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
		"owner": &graphql.Field{Type: StoreOwner},
	},
})

var StoreResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "StoreResult",
	Fields: graphql.Fields{
		"store": &graphql.Field{Type: Store},
	},
})
