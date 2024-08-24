package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetAccountConfirm() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "GetAccountCofirmFunc",
		Description: `The function generates and returns a reservation confirmation by generating a compressed zip document or folder in the desired format.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"Account": {
					Type: jsonschema.String,
					Description: "Booking ID",
				},
				"Format": {
					Type: jsonschema.String,
					Enum: []string{"0", "1", "2", "3"},
					Description: "File format. Must be '0' for PDF, '1' - HTML, '2' - Excel, '3' - Doc",
				},
			},
			Required: []string{"Account", "Format"},
		},
	}
}

