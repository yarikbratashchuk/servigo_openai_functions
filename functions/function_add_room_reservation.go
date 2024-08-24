package functions

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

//SDK Method is required an company 
func newAddRoomReservation() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "AddRoomReservationFunc",
		Description: `The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.`,
		Parameters: &jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
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
				"GuestLastName": {
					Type: jsonschema.String,
					Description: "Guest Last Name",
				},
				"GuestFirstName": {
					Type: jsonschema.String,
					Description: "Guest First Name",
				},
				"RoomTypeID": {
					Type: jsonschema.Integer,
					Description: "Room Type ID",
				},
				"PaidType": {
					Type: jsonschema.String,
					Enum: []string{"100", "200", "300"},
					Description: "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
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
				"Country": {
					Type: jsonschema.String,
					Description: "Guest Country",
				},
				"Address": {
					Type: jsonschema.String,
					Description: "Guest Address",
				},
				"Phone": {
					Type: jsonschema.String,
					Description: "Guest Phone",
				},
				"Fax": {
					Type: jsonschema.String,
					Description: "Guest Fax",
				},
				"Email": {
					Type: jsonschema.String,
					Description: "Guest Email",
				},
				"BirthDate": {
					Type: jsonschema.String,
					Description: "Guest BirthDate in format YYYY-MM-DD",
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
			Required: []string{"DateArrival", "DateDeparture", "Adults", "GuestLastName", "GuestFirstName", "RoomTypeID", "PaidType", "NeedTransport", "IsTouristTax"},
		},
	}
}

