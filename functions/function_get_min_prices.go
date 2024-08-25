package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newGetMinPrices() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_min_prices",
		Description: `The function returns the minimum possible prices by room category for the specified period.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"start_date": {
					Type: jsonschema.String,
					Description: "Start date of the search period in format YYYY-MM-DD",
				},
				"end_date": {
					Type: jsonschema.String,
					Description: "End date of the search period in format YYYY-MM-DD",
				},
			},
			Required: []string{"start_date", "end_date"},
		},
	}
}

