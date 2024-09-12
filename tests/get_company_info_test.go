package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rostis232/servigo_openai_functions/servioservice"
	"github.com/stretchr/testify/assert"
)

func TestGetCompanyInfo(t *testing.T){
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .tests_env file")
	}
	host := os.Getenv("HOST")
	accessToken := os.Getenv("ACCESS_TOKEN")

	servioService, err := servioservice.NewServioService(accessToken, host)
	assert.NoError(t, err)

	outoput, err := servioService.GetCompanyInfo("")
	t.Log(outoput)
	assert.NoError(t, err)
	assert.NotEmpty(t, outoput)
}