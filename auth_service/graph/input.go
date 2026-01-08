package schema

import (
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

var SignupInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "SignupInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"mobile_no": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"gender": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var UpdatePasswordInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdatePasswordInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"user_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"confirm_password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"current_password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var UserActivityInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserActivityInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"user_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"activity": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

