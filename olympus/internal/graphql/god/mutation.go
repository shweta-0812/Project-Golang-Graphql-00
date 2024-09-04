package god

import (
	"github.com/graphql-go/graphql"
)

/*
* Graphql root mutations
 */
var currentMaxId = 5
var GodRootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addGod": &graphql.Field{
			Type:        godType, // the return type for this field
			Description: "add a new god",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)
				description, _ := params.Args["description"].(string)
				otherNames, _ := params.Args["otherNames"].([]string)

				// figure out new id
				newID := currentMaxId + 1
				currentMaxId = currentMaxId + 1

				// perform mutation operation here
				// for e.g. create a God and save to DB.
				newGod := God{
					ID:          newID,
					Name:        name,
					Description: description,
					OtherNames:  otherNames,
				}

				GodList = append(GodList, newGod)

				// return the new God object that we supposedly save to DB
				// Note here that
				// - we are returning a `God` struct instance here
				// - we previously specified the return Type to be `godType`
				// - `God` struct maps to `godType`, as defined in `godType` ObjectConfig`
				return newGod, nil
			},
		},
		"updateGod": &graphql.Field{
			Type:        godType, // the return type for this field
			Description: "Update existing god",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				affectedGod := God{}

				// Search list for god with id
				for i := 0; i < len(GodList); i++ {
					if GodList[i].ID == id {
						if _, ok := params.Args["description"]; ok {
							GodList[i].Description = params.Args["description"].(string)
						}
						if _, ok := params.Args["name"]; ok {
							GodList[i].Name = params.Args["name"].(string)
						}
						if _, ok := params.Args["otherNames"]; ok {
							GodList[i].OtherNames = params.Args["otherNames"].([]string)
						}
						// Assign updated god so we can return it
						affectedGod = GodList[i]
						break
					}
				}
				// Return affected god
				return affectedGod, nil
			},
		},
	},
})
