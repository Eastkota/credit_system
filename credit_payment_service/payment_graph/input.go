package graph

import (
	"credit_system/core/scalar"

	"github.com/graphql-go/graphql"
)

var CreatePaymentInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreatePaymentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
		},
		"store_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
		},
		"validated_by_owner_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
		},
		"linked_credit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(scalar.UUID),
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(PaymentSubmissionStatusEnum),
		},
	},
},
)
var PaymentSubmissionStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "PaymentSubmissionStatusEnum",
	Values: graphql.EnumValueConfigMap{
		"APPROVED": &graphql.EnumValueConfig{
			Value: "approved",
		},
		"REJECTED": &graphql.EnumValueConfig{
			Value: "rejected",
		},
		"PENDING": &graphql.EnumValueConfig{
			Value: "pending",
		},
	},
})
