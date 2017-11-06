package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// TestCreateAndDeleteUser ...
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

// TestGetUsersByEmail
func TestGetUsersByEmail(t *testing.T) {
	api := New()
	assert.NotNil(t, api)

	email := "foorx2@illinois.edu"
	user, err := api.GetUsersByEmail(email)
	assert.NoError(t, err, "could not get user")

	fmt.Println(user)

}

// TestGetUser
func TestGetUser(t *testing.T) {
	api := New()
	assert.NotNil(t, api)

	user := "auth0|5a0058ec6be939112c293a65"

	result, err := api.Send(http.MethodGet, "/api/v2/users/"+user, nil)
	assert.NoError(t, err, "cou")
	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	fmt.Println(result.Status)
	fmt.Println(string(responseData))
	assert.NoError(t, err, "cou")
	assert.Equal(t, http.StatusOK, result.StatusCode)
	if result.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		assert.NoError(t, err, "cou")
	}

	res := &User{}
	json.Unmarshal(responseData, res)
	fmt.Println(res)
	assert.NoError(t, err, "cou")

	assert.NotEmpty(t, res)
}

//TestFindUser
func TestFindUser(t *testing.T) {
	api := New()
	assert.NotNil(t, api)

	username := "fooba2"

	up, err := api.getUserPage(0)
	assert.NoError(t, err)

	for pageNum := 0; pageNum < up.Length; pageNum++ {
		up, err = api.getUserPage(pageNum)
		assert.NoError(t, err)
		for _, u := range up.Users {
			if u.Username == username {
				fmt.Println(u)
			}
		}
	}
	assert.FailNow(t, "Should have found user")
}

// TestMain ...
func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}
