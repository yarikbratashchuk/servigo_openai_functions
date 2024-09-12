package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rostis232/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

//"name": "get_rooms"
const getMinPricesData = `{
	"start_date": "2024-09-01",
	"end_date": "2024-09-30"
}`

func TestGetMinPrices(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.GetMinPrices(getMinPricesData)
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}