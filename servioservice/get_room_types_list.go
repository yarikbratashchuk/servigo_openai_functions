package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) GetRoomTypesList(string) (string, error) {
	ctx := context.Background()

	companyInfo, err := s.Client.GetRoomTypesList(ctx, hotelID)
	if err != nil{
		return "", fmt.Errorf("error getting room types list: %w", err)
	}

	companyInfoBytes, err := json.Marshal(companyInfo)
	if err != nil{
		return "", fmt.Errorf("error marshaling room types list: %w", err)
	}

	return string(companyInfoBytes), nil
}