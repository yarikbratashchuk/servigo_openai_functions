package servioservice

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ServioService) GetCompanyInfo(string) (string, error) {
	ctx := context.Background()

	companyInfo, err := s.Client.GetCompanyInfo(ctx, companyCode)
	if err != nil{
		return "", fmt.Errorf("error getting company information: %w", err)
	}

	companyInfoBytes, err := json.Marshal(companyInfo)
	if err != nil{
		return "", fmt.Errorf("error marshaling company information: %w", err)
	}

	return string(companyInfoBytes), nil
}