package graph

import "github.com/graphql-go/graphql"

var PaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentResponse",
	Fields: graphql.Fields{
		"data":    &graphql.Field{Type: PaymentResult},
		"message": &graphql.Field{Type: graphql.String},
	},
})
var PaymentResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentResult",
	Fields: graphql.Fields{
		"create_payment_request": &graphql.Field{Type: CreatePaymentRequest},
	},
})
