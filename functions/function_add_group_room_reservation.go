package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newAddGroupRoomReservation() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "AddGroupRoomReservationFunc",
		Description: `The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"RoomReservations": {
					Type: jsonschema.Array,
					Description: "Array of group reservations",
					Items: &jsonschema.Definition{
						Type: jsonschema.Object,
						Properties: map[string]jsonschema.Definition{
							"GuestName": {
								Type: jsonschema.String,
								Description: "Guest Full Name",
							},
							"DateArrival": {
								Type: jsonschema.String,
								Description: "Date of arrival in format YYYY-MM-DD",
							},
							"DateDeparture": {
								Type: jsonschema.String,
								Description: "Date of departure in format YYYY-MM-DD",
							},
							"TimeArrival": {
								Type: jsonschema.String,
								Description: "Time of arrival in format HH:MM",
							},
							"TimeDeparture": {
								Type: jsonschema.String,
								Description: "Time of arrival in format HH:MM",
							},
							"Adults": {
								Type: jsonschema.Integer,
								Description: "Number of adults. Must be equal or greater than 1.",
							},
							"ChildAges": {
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{Type: jsonschema.Integer},
								Description: "Array of children's ages",
							},
							"IsExtraBedUsed": {
								Type: jsonschema.Boolean,
								Description: "Is there a need for an extra bed.",
							},
							"RoomTypeID": {
								Type: jsonschema.Integer,
								Description: "Room Type ID",
							},
							"NeedTransport": {
								Type: jsonschema.String,
								Enum: []string{"0", "1"},
								Description: "0 if the guest does not need to order transportation, 1 if the guest needs to order transportation",
							},
							"IsTouristTax": {
								Type: jsonschema.String,
								Enum: []string{"0", "1"},
								Description: "A sign of inclusion of the tourist tax in the automatically created invoice. 1 if the tourist tax is included (the guest is traveling for tourist purposes), 0 if the tourist tax is not included (the guest is traveling on business, is on a business trip).",
							},
							"AgencyCategory": {
								Type: jsonschema.String,
								Enum: []string{"0", "1", "2"},
								Description: "The guest category. 0 - not specified (the default category configured for the hotel will be used), 1 - Resident, 2 - Non-resident.",
							},
						},
					},
					Required: []string{"DateArrival", "DateDeparture", "Adults", "GuestName", "RoomTypeID", "NeedTransport", "IsTouristTax"},
				},
				"GroupName": {
					Type: jsonschema.String,
					Description: "Group Name",
				},
				"PaidType": {
					Type: jsonschema.String,
					Enum: []string{"100", "200", "300"},
					Description: "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
				},
				"Country": {
					Type: jsonschema.String,
					Description: "Guest Country",
				},
				"ContactInfo": {
					Type: jsonschema.String,
					Description: "Contact Person Info",
				},
				"ContactEmail": {
					Type: jsonschema.String,
					Description: "Contact Person Email",
				},
				"Comment": {
					Type: jsonschema.String,
					Description: "Any Comment",
				},
				"ContactName": {
					Type: jsonschema.String,
					Description: "Contact Person Name",
				},
			},
			Required: []string{"RoomReservations", "GroupName", "PaidType"},
		},
	}
}

