package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yarikbratashchuk/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

//"name": "get_rooms"
const addRoomReservationData = `{
    "address": "123 Main St, Springfield",
    "adults": 2,
    "agency_category": "1",
    "birth_date": "1990-05-12",
    "child_ages": [5, 7],
    "comment": "Early check-in requested",
    "contact_name": "John Doe",
    "country": "USA",
    "date_arrival": "2024-12-15",
    "date_departure": "2024-12-20",
    "email": "johndoe@example.com",
    "fax": "555-123456",
    "guest_first_name": "John",
    "guest_last_name": "Doe",
    "is_extra_bed_used": true,
    "is_tourist_tax": "1",
    "need_transport": "1",
    "paid_type": "200",
    "phone": "555-654321",
    "room_type_id": 2,
    "time_arrival": "14:00",
    "time_departure": "11:00"
  }`

func TestAddRoomReservation(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.AddRoomReservation(addRoomReservationData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}