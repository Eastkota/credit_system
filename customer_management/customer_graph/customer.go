package graph

import (
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

var Customer = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: scalar.UUID},
		"store_id":     &graphql.Field{Type: scalar.UUID},
		"name":         &graphql.Field{Type: graphql.String},
		"phone_number": &graphql.Field{Type: graphql.String},
		"is_active":    &graphql.Field{Type: graphql.Boolean},
	},
})

var Credit = graphql.NewObject(graphql.ObjectConfig{
	Name: "Credit",
	Fields: graphql.Fields{
		"id":                &graphql.Field{Type: scalar.UUID},
		"store_id":          &graphql.Field{Type: scalar.UUID},
		"customer_id":       &graphql.Field{Type: scalar.UUID},
		"amount":            &graphql.Field{Type: graphql.Float},
		"transaction_type":  &graphql.Field{Type: graphql.String},
		"transaction_date":  &graphql.Field{Type: graphql.String},
		"items_description": &graphql.Field{Type: graphql.String},
		"journal_number":    &graphql.Field{Type: graphql.String},
	},
})

var BalanceUpdate = graphql.NewObject(graphql.ObjectConfig{
	Name: "BalanceUpdate",
	Fields: graphql.Fields{
		"id":                      &graphql.Field{Type: scalar.UUID},
		"customer_id":             &graphql.Field{Type: scalar.UUID},
		"store_id":                &graphql.Field{Type: scalar.UUID},
		"total_credit_given":      &graphql.Field{Type: graphql.Float},
		"total_payments_received": &graphql.Field{Type: graphql.Float},
		"outstanding_balance":     &graphql.Field{Type: graphql.Float},
		"last_credit_date":        &graphql.Field{Type: graphql.String},
		"last_payment_date":       &graphql.Field{Type: graphql.String},
		"last_transaction_date":   &graphql.Field{Type: graphql.String},
		"created_at":              &graphql.Field{Type: graphql.String},
		"modified_at":             &graphql.Field{Type: graphql.String},
	},
})
