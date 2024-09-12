package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rostis232/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

//"name": "get_rooms"
const getGroupRoomsData = `{
			"booking_requests": [
				{
					"adults": 2,
					"child_ages": [5, 7],
					"date_arrival": "2024-09-15",
					"date_departure": "2024-09-20",
					"is_extra_bed_used": true,
					"time_arrival": "14:00",
					"time_departure": "11:00"
				},
				{
					"adults": 1,
					"date_arrival": "2024-09-16",
					"date_departure": "2024-09-18"
				}
			]
		}`

func TestGetGroupRooms(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.GetGroupRooms(getGroupRoomsData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}