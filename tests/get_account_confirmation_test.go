package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yarikbratashchuk/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

//"name": "get_room_types_min_price"
const getAccountConfirmationData = `{
			"account": "0000000007",
			"format": "0"
		}`

func TestGetAccountConfirmation(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.GetAccountConfirm(getAccountConfirmationData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}