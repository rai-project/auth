package api

import (
	"os"
	"testing"

	"github.com/rai-project/config"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	api := New()
	assert.NotNil(t, api)
	user, err := api.CreateUser(CreateUserRequestData{
		Username: "fooba2",
		Email:    "foorx2@illinois.edu",
		Password: "raipass2",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}
