package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetRoomType() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "GetRoomTypeFunc",
		Description: `The function returns a full description of the room by ID and localizations`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"RoomTypeID": {
					Type: jsonschema.Integer,
					Description: "Room type ID.",
				},
			},
			Required: []string{"RoomTypeID"},
		},
	}
}

