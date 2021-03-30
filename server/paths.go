package server

import (
	"github.com/graphql-go/graphql"
	"github.com/karthiklsarma/cedar-engine/m/logging"
)

var UsersList []User

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func StartGraphQlServer() graphql.Schema {
	queryFields := graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "List of users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return UsersList, nil
			},
		},
	}

	mutationFields := graphql.Fields{
		"users": &graphql.Field{
			Type:        UserType,
			Description: "Add New User",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, _ := p.Args["id"].(int)
				username, _ := p.Args["username"].(string)
				password, _ := p.Args["password"].(string)

				newUser := User{
					Id:       id,
					Username: username,
					Password: password,
				}

				UsersList = append(UsersList, newUser)
				return newUser, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root Query",
		Fields:      queryFields,
	}

	rootMutation := graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Root Mutation",
		Fields:      mutationFields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logging.Fatal("Invalid graphQl schema")
	}

	return schema
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
