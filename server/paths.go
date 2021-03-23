package server

import (
	"github.com/graphql-go/graphql"
	"github.com/karthiklsarma/cedar-engine/m/logging"
)

func StartGraphQlServer() graphql.Schema {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "cedar", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logging.Fatal("Invalid graphQl schema")
	}

	return schema
}
