package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) GetRoomTypesMinPrice(agrs string) (string, error) {
	var arguments struct {
		RoomTypeIDs []int `json:"room_type_ids"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_room_types_min_price arguments: %w", err)
	}

	ctx := context.Background()

	rooms, err := s.Client.GetRoomTypesMinPriceRequest(ctx, hotelID, arguments.RoomTypeIDs)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(rooms)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about room types min prices: %w", err)
	}

	return string(roomsBytes), nil
}