package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newCancelReservation() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "cancel_reservation",
		Description: `The function will cancel (cancel the reservation) of a guest, group or event, if possible.`,
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

