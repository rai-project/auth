package auth0

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Api struct {
	options Options
}

type Options struct {
	URL               string
	Token             string
	DefaultConnection string
}

type Option func(*Options)

func URL(s string) Option {
	return func(o *Options) {
		o.URL = s
	}
}

func Token(s string) Option {
	return func(o *Options) {
		o.Token = s
	}
}

func New(opts ...Option) *Api {
	return nil
}

func (api *Api) Send(method, endpointUrl string, body interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, api.options.URL+endpointUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.options.Token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	return client.Do(req)
}

func (api *Api) Options() Options {
	return api.options
}
