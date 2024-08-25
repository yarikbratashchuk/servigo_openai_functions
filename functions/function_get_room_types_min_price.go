package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func getRoomTypesMinPrice() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_room_types_min_price",
		Description: `The function retrieves the minimum possible price by room category for 30 days in advance.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"room_type_ids": {
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{Type: jsonschema.Integer},
					Description: "Array of room type IDs. The list of room types and their IDs can be obtained using the get_room_types_list function",
				},				
			},
			Required: []string{"room_type_ids"},
		},
	}
}

