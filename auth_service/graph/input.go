package graph

import (

	"github.com/graphql-go/graphql"
)

var SignupInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "SignupInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"phone_number": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"gender": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"store_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
