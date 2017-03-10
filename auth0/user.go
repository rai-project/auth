package auth0

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (api *Api) CreateUser(createUserRequestData CreateUserRequestData) (*CreateUserResponseData, error) {
	if len(createUserRequestData.Connection) == 0 {
		createUserRequestData.Connection = api.DefaultConnection
	}
	result, err := api.Send(http.MethodPost, "/api/v2/users", createUserRequestData)
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

	res := &CreateUserResponseData{}
	if err = json.Unmarshal(responseData, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (api *Api) GetUser(getUserRequestData GetUserRequestData) (*GetUserResponseData, error) {
	result, err := api.Send(http.MethodGet, "/api/v2/users/"+getUserRequestData.ID, nil)
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

	res := &GetUserResponseData{}
	if err = json.Unmarshal(responseData, res); err != nil {
		return nil, err
	}

	return res, nil
}

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

	if result.StatusCode != http.StatusCreated {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(responseData, errorResponse)
		if err != nil {
			return err
		}

		return errorResponse
	}

	return nil
}
