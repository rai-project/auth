package auth0

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Api struct {
	token   string
	options Options
}

type Options struct {
	Domain            string
	ClientID          string
	ClientSecret      string
	DefaultConnection string
}

type Option func(*Options)

func URL(s string) Option {
	return func(o *Options) {
		o.ClientID = s
	}
}

func ClientID(s string) Option {
	return func(o *Options) {
		o.ClientID = s
	}
}

func ClientSecret(s string) Option {
	return func(o *Options) {
		o.ClientID = s
	}
}

func New(opts ...Option) *Api {
	return nil
}

func (api *Api) getToken() error {
	type req struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Audience     string `json:"audience"`
		GrantType    string `json:"grant_type"`
	}
	url := "https://" + api.options.Domain + "/oauth/token"

}

func (api *Api) Send(method, endpointUrl string, body interface{}) (*http.Response, error) {
	if api.token == "" {
		err := api.getToken()
		if err != nil {
			return nil, err
		}
	}

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, api.options.Domain+endpointUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.options.token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return http.DefaultClient.Do(req)
}

func (api *Api) Options() Options {
	return api.options
}
