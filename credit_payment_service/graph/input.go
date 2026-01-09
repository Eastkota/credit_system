package schema

import (

	"github.com/graphql-go/graphql"
)

var CreatePaymentInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreatePaymentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"store_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"journal_number": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})