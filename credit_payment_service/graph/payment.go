package graph

var CreatePaymentRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatePaymentRequest",
	Fields: graphql.Fields{
		"customer_id": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"store_id": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"amount": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"journal_number": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"status": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"created_at": &graphql.Field{Type: scalar.DateTime},
		"updated_at": &graphql.Field{Type: scalar.DateTime},
	},
})