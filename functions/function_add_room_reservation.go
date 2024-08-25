package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newAddRoomReservation() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "add_room_reservation",
		Description: `The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.`,
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
				"guest_last_name": {
					Type: jsonschema.String,
					Description: "Guest Last Name",
				},
				"guest_first_name": {
					Type: jsonschema.String,
					Description: "Guest First Name",
				},
				"room_type_id": {
					Type: jsonschema.Integer,
					Description: "Room Type ID. The list of room types and their IDs can be obtained using the get_room_types_list function",
				},
				"paid_type": {
					Type: jsonschema.String,
					Enum: []string{"100", "200", "300"},
					Description: "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
				},
				"need_transport": {
					Type: jsonschema.String,
					Enum: []string{"0", "1"},
					Description: "0 if the guest does not need to order transportation, 1 if the guest needs to order transportation",
				},
				"is_tourist_tax": {
					Type: jsonschema.String,
					Enum: []string{"0", "1"},
					Description: "A sign of inclusion of the tourist tax in the automatically created invoice. 1 if the tourist tax is included (the guest is traveling for tourist purposes), 0 if the tourist tax is not included (the guest is traveling on business, is on a business trip).",
				},
				"agency_category": {
					Type: jsonschema.String,
					Enum: []string{"0", "1", "2"},
					Description: "The guest category. 0 - not specified (the default category configured for the hotel will be used), 1 - Resident, 2 - Non-resident.",
				},
				"country": {
					Type: jsonschema.String,
					Description: "Guest Country",
				},
				"address": {
					Type: jsonschema.String,
					Description: "Guest Address",
				},
				"phone": {
					Type: jsonschema.String,
					Description: "Guest Phone",
				},
				"fax": {
					Type: jsonschema.String,
					Description: "Guest Fax",
				},
				"email": {
					Type: jsonschema.String,
					Description: "Guest Email",
				},
				"birth_date": {
					Type: jsonschema.String,
					Description: "Guest Birth Day in format YYYY-MM-DD",
				},
				"comment": {
					Type: jsonschema.String,
					Description: "Any comment",
				},
				"contact_name": {
					Type: jsonschema.String,
					Description: "Contact Person Name",
				},
			},
			Required: []string{"date_arrival", "date_departure", "adults", "guest_last_name", "guest_first_name", "room_type_id", "paid_type", "need_transport", "is_tourist_tax"},
		},
	}
}

