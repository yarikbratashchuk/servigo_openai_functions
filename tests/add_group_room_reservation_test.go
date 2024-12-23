package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yarikbratashchuk/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

const addGroupRoomReservationData = `{
    "comment": "Special requests for the group",
    "contact_email": "contact@example.com",
    "contact_info": "1234567890",
    "contact_name": "Jane Doe",
    "country": "Ukraine",
    "group_name": "Business Trip Group",
    "paid_type": "100",
    "room_reservations": [
      {
        "adults": 2,
        "agency_category": "1",
        "child_ages": [4],
        "date_arrival": "2024-12-01",
        "date_departure": "2024-12-05",
        "guest_name": "John Smith",
        "is_extra_bed_used": false,
        "is_tourist_tax": false,
        "need_transport": "1",
        "room_type_id": 2,
        "time_arrival": "15:00",
        "time_departure": "12:00"
      }
    ]
  }`

func TestAddGroupRoomReservation(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.AddGroupRoomReservation(addGroupRoomReservationData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}