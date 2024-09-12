package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rostis232/servigo"
)

func (s *ServioService) AddGroupRoomReservation(agrs string) (string, error) {
	var arguments struct {
		Comment         string             `json:"comment,omitempty"`
		ContactEmail    string             `json:"contact_email,omitempty"`
		ContactInfo     string             `json:"contact_info,omitempty"`
		ContactName     string             `json:"contact_name,omitempty"`
		Country         string             `json:"country,omitempty"`
		GroupName       string             `json:"group_name"`
		PaidType        string             `json:"paid_type"`
		RoomReservations []struct{
			Adults          int      `json:"adults"`
			AgencyCategory  string   `json:"agency_category,omitempty"`
			ChildAges       []int    `json:"child_ages,omitempty"`
			DateArrival     string   `json:"date_arrival"`
			DateDeparture   string   `json:"date_departure"`
			GuestName       string   `json:"guest_name"`
			IsExtraBedUsed  bool     `json:"is_extra_bed_used,omitempty"`
			IsTouristTax    string   `json:"is_tourist_tax"`
			NeedTransport   string   `json:"need_transport"`
			RoomTypeID      int      `json:"room_type_id"`
			TimeArrival     string   `json:"time_arrival,omitempty"`
			TimeDeparture   string   `json:"time_departure,omitempty"`
		} `json:"room_reservations"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_rooms attributes: %w", err)
	}

	paidType, err := strconv.Atoi(arguments.PaidType)
	if err != nil {
		return "", fmt.Errorf("wrong paid_type format: %w", err)
	}

	var roomReservations []servigo.RoomReservation

	for _, roomReservationItem := range arguments.RoomReservations {

		needTransport, _ := strconv.Atoi(roomReservationItem.NeedTransport)

		var isTouristTax bool

		if roomReservationItem.IsTouristTax == "1" {
			isTouristTax = true
		}

		agencyCategory, _ := strconv.Atoi(roomReservationItem.AgencyCategory)

		roomReservation := servigo.RoomReservation{
			GuestName: roomReservationItem.GuestName,
			DateArrival: servigo.DateOnly(roomReservationItem.DateArrival),
			DateDeparture: servigo.DateOnly(roomReservationItem.DateDeparture),
			TimeArrival: servigo.TimeOnly(roomReservationItem.TimeArrival),
			TimeDeparture: servigo.TimeOnly(roomReservationItem.TimeDeparture),
			Adults: roomReservationItem.Adults,
			Childs: len(roomReservationItem.ChildAges),
			ChildAges: roomReservationItem.ChildAges,
			IsExtraBedUsed: roomReservationItem.IsExtraBedUsed,
			RoomTypeID: roomReservationItem.RoomTypeID,
			NeedTransport: needTransport,
			IsTouristTax: isTouristTax,
			AgentCategory: agencyCategory,
			ContractConditionID: contractConditionID,
			PriceListID: priceListID,
		}
		roomReservations = append(roomReservations, roomReservation)
	}

	req := servigo.AddGroupRoomReservationRequest{
		HotelID: hotelID,
		GroupName: arguments.GroupName,
		CompanyID: companyID,
		Country: arguments.Country,
		PaidType: servigo.PaidType(paidType),
		ContactName: arguments.ContactName,
		ContactEMail: arguments.ContactEmail,
		ContactInfo: arguments.ContactInfo,
		Comment: arguments.Comment,
		ContractConditionID: contractConditionID,
		RoomReservations: roomReservations,
	}

	ctx := context.Background()

	account, _, err := s.Client.AddGroupRoomReservation(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(account)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about rooms: %w", err)
	}

	return string(roomsBytes), nil
}