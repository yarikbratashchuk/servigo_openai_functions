package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) GetRoomType(agrs string) (string, error) {
	var arguments struct {
		RoomTypeID int `json:"room_type_id"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_room_type arguments: %w", err)
	}

	ctx := context.Background()

	rooms, err := s.Client.GetRoomType(ctx, arguments.RoomTypeID)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	roomsBytes, err := json.Marshal(rooms)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about get_room_type: %w", err)
	}

	return string(roomsBytes), nil
}