package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newAddGroupRoomReservation() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "add_group_room_reservation",
		Description: `The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"room_reservations": {
					Type: jsonschema.Array,
					Description: "Array of group reservations",
					Items: &jsonschema.Definition{
						Type: jsonschema.Object,
						Properties: map[string]jsonschema.Definition{
							"guest_name": {
								Type: jsonschema.String,
								Description: "Guest Full Name",
							},
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
							"room_type_id": {
								Type: jsonschema.Integer,
								Description: "Room Type ID. The list of room types and their IDs can be obtained using the get_room_types_list function",
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
						},
					},
					Required: []string{"date_arrival", "date_departure", "adults", "guest_name", "room_type_id", "need_transport", "is_tourist_tax"},
				},
				"group_name": {
					Type: jsonschema.String,
					Description: "Group Name",
				},
				"paid_type": {
					Type: jsonschema.String,
					Enum: []string{"100", "200", "300"},
					Description: "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
				},
				"country": {
					Type: jsonschema.String,
					Description: "Guest Country",
				},
				"contact_info": {
					Type: jsonschema.String,
					Description: "Contact Person Info",
				},
				"contact_email": {
					Type: jsonschema.String,
					Description: "Contact Person Email",
				},
				"comment": {
					Type: jsonschema.String,
					Description: "Any Comment",
				},
				"contact_name": {
					Type: jsonschema.String,
					Description: "Contact Person Name",
				},
			},
			Required: []string{"room_reservations", "group_name", "paid_type"},
		},
	}
}

