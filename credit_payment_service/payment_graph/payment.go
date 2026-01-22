package graph

import (
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

var CreatePaymentRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatePaymentRequest",
	Fields: graphql.Fields{
		"customer_id":           &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"store_id":              &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"amount":                &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"status":                &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"validated_by_owner_id": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"linked_credit_id":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
	},
})

var OwnerPaymentRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "OwnerPaymentRequest",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.NewNonNull(scalar.UUID)},
		"store_id":       &graphql.Field{Type: graphql.NewNonNull(scalar.UUID)},
		"screenshot_url": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"status":         &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
	},
})
