package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetRooms() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_rooms",
		Description: `The function gets the list of room categories with the number of available rooms of each category according to the input parameters.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"date_arrival": {
					Type: jsonschema.String,
					Description: "Date of arrival in format YYYY-MM-DD",
				},
				"date_departure": {
					Type: jsonschema.String,
					Description: "Date of departure in format YYYY-MM-DD",
				},
				"time_arrival": {
					Type: jsonschema.String,
					Description: "Time of arrival in format HH:MM",
				},
				"time_departure": {
					Type: jsonschema.String,
					Description: "Time of arrival in format HH:MM",
				},
				"adults": {
					Type: jsonschema.Integer,
					Description: "Number of adults. Must be equal or greater than 1.",
				},
				"child_ages": {
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{Type: jsonschema.Integer},
					Description: "Array of children's ages",
				},
				"is_extra_bed_used": {
					Type: jsonschema.Boolean,
					Description: "Is there a need for an extra bed.",
				},
			},
			Required: []string{"date_arrival", "date_departure", "adults"},
		},
	}
}

