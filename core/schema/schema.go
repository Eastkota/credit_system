package schema

import (
	"errors"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	globalSchema *graphql.Schema
	initErr      error
	once         sync.Once
)

// InitSchema is called once in main.go after merging all module fields.
func InitSchema(q *graphql.Object, m *graphql.Object) error {
	once.Do(func() {
		s, err := graphql.NewSchema(graphql.SchemaConfig{
			Query:    q,
			Mutation: m,
		})
		if err != nil {
			initErr = err
			return
		}
		globalSchema = &s
	})
	return initErr
}

// GetSchema provides the compiled schema to the handlers safely.
func GetSchema() (*graphql.Schema, error) {
	if globalSchema == nil {
		if initErr != nil {
			return nil, initErr
		}
		return nil, errors.New("schema has not been initialized")
	}
	return globalSchema, nil
}
