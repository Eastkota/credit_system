package graphqlmodule

import "github.com/graphql-go/graphql"

type Module interface {
	Queries() graphql.Fields
	Mutations() graphql.Fields
}