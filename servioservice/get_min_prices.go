package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (s *ServioService) GetMinPrices(agrs string) (string, error) {
	var arguments struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	err := json.Unmarshal([]byte(agrs), &arguments)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling get_min_prices attributes: %w", err)
	}

	ctx := context.Background()

	startDate, err := time.Parse("2006-01-02", arguments.StartDate)
	if err != nil {
		return "", fmt.Errorf("wrong start_date format: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", arguments.EndDate)
	if err != nil {
		return "", fmt.Errorf("wrong start_date format: %w", err)
	}

	minprices, err := s.Client.GetMinPrices(ctx, []int{hotelID},startDate, endDate, companyCode)
	if err != nil {
		return "", fmt.Errorf("error making request to hotel managment system: %w", err)
	}

	minpricesBytes, err := json.Marshal(minprices)
	if err != nil{
		return "", fmt.Errorf("error marshaling information about min prices: %w", err)
	}

	return string(minpricesBytes), nil
}