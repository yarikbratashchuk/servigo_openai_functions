package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) CancelReservation(agrs string) (string, error) {
	var arguments struct {
		Account string `json:"account"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling cancel_reservation arguments: %w", err)
	}

	ctx := context.Background()

	err = s.Client.CancelReservation(ctx, arguments.Account)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	reservationInfoBytes, err := json.Marshal("reservation successfuly canceled")
	if err != nil{
		return "", fmt.Errorf("error marshaling information about cancel_reservation: %w", err)
	}

	return string(reservationInfoBytes), nil
}