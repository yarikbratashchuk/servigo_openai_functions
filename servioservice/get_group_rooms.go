package servioservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yarikbratashchuk/servigo"
)

func (s *ServioService) GetGroupRooms(agrs string) (string, error) {
	var arguments struct {
		BookingRequests []struct{
			Adults         int    `json:"adults"`
			ChildAges      []int  `json:"child_ages,omitempty"`
			DateArrival    string `json:"date_arrival"`
			DateDeparture  string `json:"date_departure"`
			IsExtraBedUsed bool   `json:"is_extra_bed_used,omitempty"`
			TimeArrival    string `json:"time_arrival,omitempty"`
			TimeDeparture  string `json:"time_departure,omitempty"`
		} `json:"booking_requests"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_rooms attributes: %w", err)
	}

	var bookingRequests []servigo.BookingRequest

	for _, brArgumentItem := range arguments.BookingRequests {
		bookingRequest, err := servigo.NewBookingRequest(servigo.DateOnly(brArgumentItem.DateArrival), 
			servigo.DateOnly(brArgumentItem.DateDeparture), 
			servigo.TimeOnly(brArgumentItem.TimeArrival),
			servigo.TimeOnly(brArgumentItem.TimeDeparture), 
			brArgumentItem.Adults, brArgumentItem.ChildAges, brArgumentItem.IsExtraBedUsed)
		if err != nil {
			return "", fmt.Errorf("error creating request with these attributes: %w", err)
		}

		bookingRequests = append(bookingRequests, *bookingRequest)
	}

	req, err := servigo.NewGetGroupRoomsRequest(hotelID, bookingRequests, companyCodeID, servigo.Ukrainian)
		if err != nil {
			return "", fmt.Errorf("error creating request with these attributes: %w", err)
		}

	ctx := context.Background()

	rooms, _, err := s.Client.GetGroupRooms(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(rooms)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about rooms: %w", err)
	}

	return string(roomsBytes), nil
}