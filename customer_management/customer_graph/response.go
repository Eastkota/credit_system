package graph

import "github.com/graphql-go/graphql"

var CustomerResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CustomerResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var CustomerResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerResult",
	Fields: graphql.Fields{
		"customer": &graphql.Field{Type: Customer},
	},
})
