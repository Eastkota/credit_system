package graph

import (
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

var CustomerInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CustomerInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"store_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"phone_number": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var CreditInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreditInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"store_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"customer_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"transaction_type": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(TransactionTypeEnum),
			},
			"items_description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"journal_number": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var TransactionTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "TransactionType",
	Values: graphql.EnumValueConfigMap{
		"CREDIT_GIVEN": &graphql.EnumValueConfig{
			Value: "credit_given",
		},
		"PAYMENTS_RECEIVED": &graphql.EnumValueConfig{
			Value: "payment_received",
		},
	},
})
