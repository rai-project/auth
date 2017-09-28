package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Api ...
type Api struct {
	token   string
	options Options
}

// Options ...
type Options struct {
	Domain       string
	ClientID     string
	ClientSecret string
	Connection   string
}

// Option ...
type Option func(*Options)

// Domain ...
func Domain(s string) Option {
	return func(o *Options) {
		o.Domain = strings.TrimLeft(s, "https://")
	}
}

// ClientID ...
func ClientID(s string) Option {
	return func(o *Options) {
		o.ClientID = s
	}
}

// ClientSecret ...
func ClientSecret(s string) Option {
	return func(o *Options) {
		o.ClientSecret = s
	}
}

// Connection ...
func Connection(s string) Option {
	return func(o *Options) {
		o.Connection = s
	}
}

// New ...
func New(iopts ...Option) *Api {
	opts := Options{
		Domain:       strings.TrimLeft(Config.Domain, "https://"),
		ClientID:     Config.ClientID,
		ClientSecret: Config.ClientSecret,
		Connection:   Config.Connection,
	}
	for _, o := range iopts {
		o(&opts)
	}
	return &Api{options: opts}
}

func (api *Api) getToken() error {
	type q struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Audience     string `json:"audience"`
		GrantType    string `json:"grant_type"`
	}
	type r struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	url := "https://" + api.options.Domain + "/oauth/token"
	audience := "https://" + api.options.Domain + "/api/v2/"

	request := q{
		ClientID:     api.options.ClientID,
		ClientSecret: api.options.ClientSecret,
		Audience:     audience,
		GrantType:    "client_credentials",
	}

	bts, err := json.Marshal(request)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(bts)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var t r
	if err := json.Unmarshal(body, &t); err != nil {
		return err
	}

	api.token = t.AccessToken

	return nil
}

// Send ...
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

	req, err := http.NewRequest(method, "https://"+api.options.Domain+endpointUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return http.DefaultClient.Do(req)
}

// Options ...
func (api *Api) Options() Options {
	return api.options
}
