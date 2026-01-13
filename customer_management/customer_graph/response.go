package graph

import "github.com/graphql-go/graphql"

var CustomerResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CustomerResult},
		"error": &graphql.Field{Type: CustomerError},
	},
})

var CreditResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreditResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CreditResult},
		"error": &graphql.Field{Type: CustomerError},
	},
})

var CustomerResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "CustomerResult",
	Fields: graphql.Fields{
		"customer": &graphql.Field{Type: Customer},
	},
})

var CreditResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreditResult",
	Fields: graphql.Fields{
		"credit": &graphql.Field{Type: Credit},
	},
})

var BalanceResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "BalanceResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: BalanceResult},
		"error": &graphql.Field{Type: CustomerError},
	},
})

var BalanceResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "BalanceResult",
	Fields: graphql.Fields{
		"balance": &graphql.Field{Type: BalanceUpdate},
	},
})

var MultipleBalanceResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "MultipleBalanceResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: MultipleBalanceResult},
		"error": &graphql.Field{Type: CustomerError},
	},
})

var MultipleBalanceResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "MultipleBalanceResult",
	Fields: graphql.Fields{
		"balances": &graphql.Field{Type: graphql.NewList(BalanceUpdate)},
	},
})
