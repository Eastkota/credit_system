package graph

import "github.com/graphql-go/graphql"

var CustomerResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: Customer},
		"error": &graphql.Field{Type: AuthError},
	},
})
