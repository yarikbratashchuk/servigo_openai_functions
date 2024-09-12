package servioservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rostis232/servigo"
)

func (s *ServioService) GetRooms(agrs string) (string, error) {
	var arguments struct {
		Adults         int    `json:"adults"`
		ChildAges      []int  `json:"child_ages,omitempty"`
		DateArrival    string `json:"date_arrival"`
		DateDeparture  string `json:"date_departure"`
		IsExtraBedUsed bool   `json:"is_extra_bed_used,omitempty"`
		TimeArrival    string `json:"time_arrival,omitempty"`
		TimeDeparture  string `json:"time_departure,omitempty"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_rooms attributes: %w", err)
	}

	bookingRequest, err := servigo.NewBookingRequest(servigo.DateOnly(arguments.DateArrival), 
		servigo.DateOnly(arguments.DateDeparture), 
		servigo.TimeOnly(arguments.TimeArrival),
		servigo.TimeOnly(arguments.TimeDeparture), 
		arguments.Adults, arguments.ChildAges, arguments.IsExtraBedUsed)
		if err != nil {
			return "", fmt.Errorf("error creating request with these attributes: %w", err)
		}
	
	req, err := servigo.NewGetRoomsRequest(hotelID, *bookingRequest, companyCodeID, servigo.Ukrainian)
	if err != nil {
		return "", fmt.Errorf("error creating request with these attributes: %w", err)
	}

	ctx := context.Background()

	rooms, _, err := s.Client.GetRooms(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(rooms)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about rooms: %w", err)
	}

	return string(roomsBytes), nil
}