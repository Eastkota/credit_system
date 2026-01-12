package schema

import (
	"fmt"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	queryType     *graphql.Object
	mutationType  *graphql.Object
	schema        graphql.Schema
	schemaErr     error
	once          sync.Once
	isInitialized bool
)

func InitSchema(q *graphql.Object, m *graphql.Object) error {
	var initErr error

	once.Do(func() {
		if q == nil {
			initErr = fmt.Errorf("query type cannot be nil")
			schemaErr = initErr
			return
		}

		queryType = q
		mutationType = m

		config := graphql.SchemaConfig{
			Query: queryType,
		}

		// Only add mutation if provided
		if mutationType != nil {
			config.Mutation = mutationType
		}

		schema, schemaErr = graphql.NewSchema(config)
		if schemaErr != nil {
			initErr = fmt.Errorf("failed to create schema: %w", schemaErr)
			return
		}

		isInitialized = true
		fmt.Println("✅ Schema initialized successfully")
	})

	if initErr != nil {
		return initErr
	}

	return schemaErr
}

func GetSchema() (*graphql.Schema, error) {
	if !isInitialized {
		return nil, fmt.Errorf("schema not initialized - call InitSchema first")
	}

	if schemaErr != nil {
		return nil, schemaErr
	}

	return &schema, nil
}

func IsInitialized() bool {
	return isInitialized
}

func ResetSchema() {
	once = sync.Once{}
	isInitialized = false
	schemaErr = nil
}
