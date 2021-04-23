package server

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/karthiklsarma/cedar-server/m/logging"
	"github.com/karthiklsarma/cedar-server/m/stream"
)

func StartGraphQlServer() graphql.Schema {
	queryFields := graphql.Fields{
		"getUsers": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "List of users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return UsersList, nil
			},
		},
		"getLocations": &graphql.Field{
			Type:        graphql.NewList(LocationType),
			Description: "Location of users",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"group": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username, ok := p.Args["username"].(string)
				if !ok {
					return nil, errors.New("username not provided or invalid username")
				}

				logging.Debug(fmt.Sprintf("will query locations for user: %s", username))
				return LocationList, nil
			},
		},
	}

	mutationFields := graphql.Fields{
		"setUsers": &graphql.Field{
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

				newUser := User{
					Id:       id,
					Username: username,
				}

				UsersList = append(UsersList, newUser)
				return newUser, nil
			},
		},
		"setLocation": &graphql.Field{
			Type:        LocationType,
			Description: "location of the user",
			Args: graphql.FieldConfigArgument{
				"lat": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"lng": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				lat := p.Args["lat"].(float64)
				logging.Info(fmt.Sprintf("Received lat: %v", lat))
				lng := p.Args["lng"].(float64)
				locationmsg := fmt.Sprintf(stream.LOCATION_TEMPLATE, lat, lng)

				stream.EmitLocation(locationmsg)
				logging.Info(fmt.Sprintf("Received lng: %v", lng))
				loc := &Location{}
				logging.Info(fmt.Sprintf("sending location message : %v to eventqueue", locationmsg))
				if err := json.Unmarshal([]byte(locationmsg), loc); err != nil {
					return nil, err
				}
				return loc, nil
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
