package graphql

import (
	"github.com/graphql-go/graphql"
	"olympus/internal/graphql/god"
)

var GodRootQuery = god.GodRootQuery
var GodRootMutation = god.GodRootMutation

// define Graphql Main schema, with our rootQuery and rootMutation
var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    GodRootQuery,
	Mutation: GodRootMutation,
})
