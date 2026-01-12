package graph

import "github.com/graphql-go/graphql"

var PaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentResponse",
	Fields: graphql.Fields{
		"credit_id": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"message": &graphql.Field{Type: graphql.String},
	},
})





