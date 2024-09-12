package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rostis232/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

//"name": "get_room_types_min_price"
const getRoomTypesMinPricesData = `{
			"room_type_ids": [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
		}`

func TestGetRoomTypesMinPrices(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.GetRoomTypesMinPrice(getRoomTypesMinPricesData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}