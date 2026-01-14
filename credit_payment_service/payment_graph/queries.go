package graph

import (
	"credit_system/credit_payment_service/helpers"
	"credit_system/credit_payment_service/model"
	"credit_system/credit_payment_service/payment_resolvers"
	"credit_system/core/schema"

	"github.com/graphql-go/graphql"
)

func NewPaymentQueryType(resolver *resolvers.PaymentResolver) graphql.Fields {
    return graphql.Fields{
		"payment_service": &graphql.Field{
			Type: graphql.NewNonNull(PaymentService),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				s, err := schema.GetSchema()
				if err != nil {
					return nil, err
				}

				return model.PaymentService{
					Name:    "PaymentService",
					Version: "1.0.0",
					Schema:  helpers.ConvertSchemaToString(s),
				}, nil
			},
		},
	}
}
