package graph

import (
	"github.com/graphql-go/graphql"
)

var CustomerInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CustomerInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"store_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"phone_number": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"has_credit": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)
