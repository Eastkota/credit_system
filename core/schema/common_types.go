package schema

import "github.com/graphql-go/graphql"

var ServiceInfoType = graphql.NewObject(graphql.ObjectConfig{
    Name: "ServiceInfo", 
    Fields: graphql.Fields{
        "name":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "version": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "schema":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
    },
})