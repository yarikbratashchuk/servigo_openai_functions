package servioservice

import (
	"github.com/rostis232/servigo"
)

const (
	companyCode = "QVLSXA44J%"
	companyCodeID = 1
	hotelID = 1
	companyID = 4
	contractConditionID = 1
	priceListID1 = 1
	priceListID2 = 2
)

var (
	host string
	AccessToken string
)

type ServioService struct{
	Client *servigo.Client
}

func NewServioService(accessToken, host string) (*ServioService, error) {
	servigo.SetServioHost(host)

	client, err := servigo.NewClient(accessToken)
	if err != nil {
		return nil, err
	}

	return &ServioService{
		Client: client,
	}, nil
}

