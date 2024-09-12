package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) GetReservationInfo(agrs string) (string, error) {
	var arguments struct {
		Account string `json:"account"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_reservation_type arguments: %w", err)
	}

	ctx := context.Background()

	reservationInfo, err := s.Client.GetReservationInfo(ctx, arguments.Account)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	reservationInfoBytes, err := json.Marshal(reservationInfo)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about get_reservation_type: %w", err)
	}

	return string(reservationInfoBytes), nil
}