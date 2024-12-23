package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yarikbratashchuk/servigo"
)

func (s *ServioService) AddRoomReservation(agrs string) (string, error) {
	var arguments struct {
		Address         string  `json:"address,omitempty"`
		Adults          int     `json:"adults"`
		AgencyCategory  string  `json:"agency_category,omitempty"`
		BirthDate       string  `json:"birth_date,omitempty"`
		ChildAges       []int   `json:"child_ages,omitempty"`
		Comment         string  `json:"comment,omitempty"`
		ContactName     string  `json:"contact_name,omitempty"`
		Country         string  `json:"country,omitempty"`
		DateArrival     string  `json:"date_arrival"`
		DateDeparture   string  `json:"date_departure"`
		Email           string  `json:"email,omitempty"`
		Fax             string  `json:"fax,omitempty"`
		GuestFirstName  string  `json:"guest_first_name"`
		GuestLastName   string  `json:"guest_last_name"`
		IsExtraBedUsed  bool    `json:"is_extra_bed_used,omitempty"`
		IsTouristTax    string  `json:"is_tourist_tax"`
		NeedTransport   string  `json:"need_transport"`
		PaidType        string  `json:"paid_type"`
		Phone           string  `json:"phone,omitempty"`
		RoomTypeID      int     `json:"room_type_id"`
		TimeArrival     string  `json:"time_arrival,omitempty"`
		TimeDeparture   string  `json:"time_departure,omitempty"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_rooms attributes: %w", err)
	}

	paidType, _ := strconv.Atoi(arguments.PaidType)

	agencyCategory, _ := strconv.Atoi(arguments.AgencyCategory)

	isTouristTax, err := strconv.Atoi(arguments.IsTouristTax)
	if err != nil {
		return "", fmt.Errorf("wrong is_tourist_tax format: %w", err)
	}

	needTransport, _ := strconv.Atoi(arguments.NeedTransport)

	req := servigo.AddRoomReservationRequest{
		HotelID: hotelID,
		DateArrival: servigo.DateOnly(arguments.DateArrival),
		DateDeparture: servigo.DateOnly(arguments.DateDeparture),
		TimeArrival: servigo.TimeOnly(arguments.TimeArrival),
		TimeDeparture: servigo.TimeOnly(arguments.TimeDeparture),
		Adults: arguments.Adults,
		Childs: len(arguments.ChildAges),
		ChildAges: arguments.ChildAges,
		IsExtraBedUsed: arguments.IsExtraBedUsed,
		GuestLastName: arguments.GuestLastName,
		GuestFirstName: arguments.GuestFirstName,
		RoomTypeID: arguments.RoomTypeID,
		CompanyID: companyID,
		ContractConditionID: contractConditionID,
		PaidType: servigo.PaidType(paidType),
		Address: arguments.Address,
		AgencyCategory: agencyCategory,
		BirthDate: servigo.DateOnly(arguments.BirthDate),
		Comment: arguments.Comment,
		ContactName: arguments.ContactName,
		Country: arguments.Country,
		Email: arguments.Email,
		Fax: arguments.Fax,
		IsTouristTax: isTouristTax,
		NeedTransport: needTransport,
		Phone: arguments.Phone,
		PriceListID: priceListID1,
	}

	ctx := context.Background()

	reservation, err := s.Client.AddRoomReservation(ctx,&req)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(reservation)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about rooms: %w", err)
	}

	return string(roomsBytes), nil
}