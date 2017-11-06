package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func (api *Api) getUserPage(ipage int) (*UserPage, error) {
	page := strconv.Itoa(ipage)
	result, err := api.Send(http.MethodGet, "/api/v2/users?per_page=100&page="+page+"&include_totals=true&search_engine=v2", nil)
	if err != nil {
		return nil, err
	}
	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	res := &UserPage{}
	if err = json.Unmarshal(responseData, res); err != nil {
		return nil, err
	}

	return res, nil
}

// FindUser ...
func (api *Api) FindUser(username string) (User, error) {
	if username == "" {
		return User{}, errors.New("username cannot be empty while attempting to find user")
	}
	up, err := api.getUserPage(0)
	if err != nil {
		return User{}, err
	}
	for pageNum := 0; pageNum < up.Length; pageNum++ {
		up, err = api.getUserPage(pageNum)
		if err != nil {
			return User{}, err
		}
		for _, u := range up.Users {
			if u.Username == username {
				return u, nil
			}
		}
	}
	return User{}, errors.Errorf("unable to find the user %v", username)
}

// CreateUser ...
func (api *Api) CreateUser(createUserRequestData CreateUserRequestData) (User, error) {
	if len(createUserRequestData.Connection) == 0 {
		createUserRequestData.Connection = api.options.Connection
	}
	if createUserRequestData.GivenName != "" {
		createUserRequestData.UserMetadata["name"] = createUserRequestData.GivenName
	}
	result, err := api.Send(http.MethodPost, "/api/v2/users", createUserRequestData)
	if err != nil {
		return User{}, err
	}
	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return User{}, err
	}

	if result.StatusCode == http.StatusBadRequest {
		user, err := api.FindUser(createUserRequestData.Username)
		if err != nil {
			return User{}, err
		}

		baseRequest := UpdateUserRequestData{
			Connection:    createUserRequestData.Connection,
			AppMetadata:   createUserRequestData.AppMetadata,
			PhoneNumber:   createUserRequestData.PhoneNumber,
			PhoneVerified: createUserRequestData.PhoneVerified,
			UserMetadata:  createUserRequestData.UserMetadata,
		}
		if createUserRequestData.GivenName != "" {
			baseRequest.UserMetadata["name"] = createUserRequestData.GivenName
		}

		if createUserRequestData.Email != "" {
			req := baseRequest
			req.Email = createUserRequestData.Email
			user, err = api.UpdateUser(user.UserID, req)
			if err != nil {
				return User{}, err
			}
		}
		if createUserRequestData.VerifyEmail {
			req := baseRequest
			req.VerifyEmail = createUserRequestData.VerifyEmail
			user, err = api.UpdateUser(user.UserID, req)
			if err != nil {
				return User{}, err
			}
		}
		if createUserRequestData.EmailVerified {
			req := baseRequest
			req.EmailVerified = createUserRequestData.EmailVerified
			user, err = api.UpdateUser(user.UserID, req)
			if err != nil {
				return User{}, err
			}
		}
		if createUserRequestData.Password != "" {
			req := baseRequest
			req.Password = createUserRequestData.Password
			user, err = api.UpdateUser(user.UserID, req)
			if err != nil {
				return User{}, err
			}
		}
		return user, nil
	}

	if result.StatusCode != http.StatusOK && result.StatusCode != http.StatusCreated {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return User{}, err
		}

		return User{}, errorResponse
	}

	res := User{}
	if err = json.Unmarshal(responseData, &res); err != nil {
		return User{}, err
	}

	return res, nil
}

// UpdateUser ...
func (api *Api) UpdateUser(userID string, updateUserRequestData UpdateUserRequestData) (User, error) {
	if len(updateUserRequestData.Connection) == 0 {
		updateUserRequestData.Connection = api.options.Connection
	}
	result, err := api.Send(http.MethodPatch, "/api/v2/users/"+userID, updateUserRequestData)
	if err != nil {
		return User{}, err
	}

	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return User{}, err
	}

	if result.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return User{}, err
		}

		return User{}, errorResponse
	}

	res := User{}
	if err = json.Unmarshal(responseData, &res); err != nil {
		return User{}, err
	}

	return res, nil
}

// DeleteUser ...
func (api *Api) DeleteUser(userID string) error {

	result, err := api.Send(http.MethodDelete, "/api/v2/users/"+userID, nil)
	if err != nil {
		return err
	}

	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	if result.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return err
		}

		return errors.New(errorResponse.Message)
	}

	return nil
}

// GetUser ...
func (api *Api) GetUser(getUserRequestData GetUserRequestData) (User, error) {
	result, err := api.Send(http.MethodGet, "/api/v2/users/"+getUserRequestData.UserID, nil)
	if err != nil {
		return User{}, err
	}

	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return User{}, err
	}

	if result.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return User{}, err
		}

		return User{}, errorResponse
	}

	res := User{}
	if err = json.Unmarshal(responseData, &res); err != nil {
		return User{}, err
	}

	return res, nil
}

// SendVerificationEmail ...
func (api *Api) SendVerificationEmail(requestData SendVerificationEmailRequestData) error {
	result, err := api.Send(http.MethodPost, "/api/v2/jobs/post_verification_email", requestData)
	if err != nil {
		return err
	}

	if result.Body != nil {
		defer result.Body.Close()
	}
	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	if result.StatusCode != http.StatusOK && result.StatusCode != http.StatusCreated {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, errorResponse)
		if err != nil {
			return err
		}

		return errorResponse
	}

	return nil
}
