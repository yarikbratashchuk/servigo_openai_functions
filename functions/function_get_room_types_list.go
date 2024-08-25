package functions

import (
	"github.com/sashabaranov/go-openai"
)

//SDK Method is required an company 
func newGetRoomTypesList() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_room_types_list",
		Description: `The function returns a list of room categories by hotel. `,
	}
}