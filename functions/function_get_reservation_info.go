package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetReservationInfo() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_reservation_info",
		Description: `The function returns information on the created reservation, data is returned only for confirmed reservations.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"account": {
					Type: jsonschema.String,
					Description: "Booking ID",
				},
			},
			Required: []string{"account"},
		},
	}
}

