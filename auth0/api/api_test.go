package api

import (
	"os"
	"testing"

	"github.com/rai-project/config"
	"github.com/stretchr/testify/assert"
)

// TestCreateUser ...
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

// TestCreateUser ...
func TestCreateAndDeleteUser(t *testing.T) {
	api := New()
	assert.NotNil(t, api)
	user, err := api.CreateUser(CreateUserRequestData{
		Username: "deleteme",
		Email:    "deleteme@illinois.edu",
		Password: "deletemepass",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	err = api.DeleteUser(user.UserID)
	assert.NoError(t, err)
}

// TestMain ...
func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}
