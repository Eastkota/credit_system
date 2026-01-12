package graph

import (
	"github.com/graphql-go/graphql"
)

var Customer = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"store_id":     &graphql.Field{Type: graphql.String},
		"name":         &graphql.Field{Type: graphql.String},
		"phone_number": &graphql.Field{Type: graphql.String},
		"has_credit":   &graphql.Field{Type: graphql.Boolean},
	},
})
