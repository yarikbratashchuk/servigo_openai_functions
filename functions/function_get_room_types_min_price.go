package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func getRoomTypesMinPrice() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "GetRoomTypesMinPrice",
		Description: `The function retrieves the minimum possible price by room category for 30 days in advance.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"RoomTypeIDs": {
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{Type: jsonschema.Integer},
					Description: "Array of room type IDs.",
				},				
			},
			Required: []string{"RoomTypeIDs"},
		},
	}
}

