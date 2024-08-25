package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetRoomType() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_room_type",
		Description: `The function returns a full description of the room by ID and localizations`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"room_type_id": {
					Type: jsonschema.Integer,
					Description: "Room type ID. The list of room types and their IDs can be obtained using the get_room_types_list function",
				},
			},
			Required: []string{"room_type_id"},
		},
	}
}

