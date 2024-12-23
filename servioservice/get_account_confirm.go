package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yarikbratashchuk/servigo"
)

func (s *ServioService) GetAccountConfirm(agrs string) (string, error) {
	var arguments struct {
		Account string `json:"account"`
		Format  string `json:"format"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_reservation_type arguments: %w", err)
	}

	fileFormatInt, err := strconv.Atoi(arguments.Format)
	if err != nil {
		return "", fmt.Errorf("wrong format value: %w", err)
	}

	req, err := servigo.NewGetAccountConfirmRequest(arguments.Account, servigo.Ukrainian, servigo.FileFormat(fileFormatInt))
	if err != nil {
		return "", fmt.Errorf("wrong creating request for getting account confirmation: %w", err)
	}

	ctx := context.Background()

	reservationInfo, err := s.Client.GetAccountConfirm(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	reservationInfoBytes, err := json.Marshal(reservationInfo)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about get_reservation_type: %w", err)
	}

	return string(reservationInfoBytes), nil
}