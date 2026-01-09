package schema

import "github.com/graphql-go/graphql"

func MergeFields(fieldMaps ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}

	for _, fields := range fieldMaps {
		for name, field := range fields {
			if _, exists := merged[name]; exists {
				panic("duplicate GraphQL field: " + name)
			}
			merged[name] = field
		}
	}

	return merged
}
