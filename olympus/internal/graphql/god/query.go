package god

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
)

// Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}

var GodList []God
var _ = importJSONDataFromFile("./sampleData.json", &GodList)

type God struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OtherNames  []string `json:"otherNames"`
}

// define custom GraphQL ObjectType `godType` for our Golang struct `God`
// - the field type matches the field type in our struct
var godType = graphql.NewObject(graphql.ObjectConfig{
	Name: "God",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"otherNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

/*
* Graphql root query
 */
var GodRootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// query 1 fields
		"god": &graphql.Field{
			Description: "Get single god",
			Type:        godType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				nameQuery, isOK := params.Args["name"].(string)
				if isOK {
					// Search for el with name
					for _, god := range GodList {
						if god.Name == nameQuery {
							return god, nil
						}
					}
				}

				return God{}, nil
			},
		},
		// query 2 fields
		"godList": &graphql.Field{
			Description: "List of gods",
			Type:        graphql.NewList(godType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GodList, nil
			},
		},
	},
})
