package graph

import "github.com/graphql-go/graphql"

var CustomerError = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
