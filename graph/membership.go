package schema

import (
	"auth_service/graph/scalar"
	
	"github.com/graphql-go/graphql"
)

var AuthUserMembership = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUserMembership",
	Fields: graphql.Fields{
		"id":                     &graphql.Field{Type: scalar.UUID},
		"membership_join_date":   &graphql.Field{Type: scalar.Time},
		"membership_end_date":    &graphql.Field{Type: scalar.Time},
		"user_id":             &graphql.Field{Type: scalar.UUID},
		"membership_duration_id": &graphql.Field{Type: scalar.UUID},
		"membership_duration": &graphql.Field{Type: AuthMembershipDuration},

	},
})

var AuthMembershipDuration = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthMembershipDuration",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: scalar.UUID},
		"duration_type":  &graphql.Field{Type: graphql.String},
		"price":          &graphql.Field{Type: graphql.String},
		"description":    &graphql.Field{Type: graphql.String},
		"package_id":     &graphql.Field{Type: scalar.UUID},
		"number_of_days": &graphql.Field{Type: graphql.String},
		"price_usd":      &graphql.Field{Type: graphql.String},
		"package":		  &graphql.Field{Type: AuthPackages},
	},
})

var AuthPackages = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthPackages",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: scalar.UUID},
		"name":         &graphql.Field{Type: graphql.String},
		"description":  &graphql.Field{Type: graphql.String},
		"package_type": &graphql.Field{Type: graphql.String},
		"created_at":   &graphql.Field{Type: scalar.Time},
		"updated_at":   &graphql.Field{Type: scalar.Time},
		"user_limit":   &graphql.Field{Type: graphql.Int},
	},
})
