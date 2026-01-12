package graph

import (

	"github.com/graphql-go/graphql"
	"credit_system/core/scalar"
)

var CreatePaymentInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreatePaymentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"store_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
		},
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
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