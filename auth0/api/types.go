package api

import "fmt"

type CreateUserRequestData struct {
	Connection    string                 `json:"connection"`
	Email         string                 `json:"email,omitempty"`
	Username      string                 `json:"username,omitempty"`
	Password      string                 `json:"password,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	VerifyEmail   bool                   `json:"verify_email,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
}

type CreateUserResponseData struct {
	Email         string                 `json:"email"`
	EmailVerified bool                   `json:"email_verified"`
	Username      string                 `json:"username"`
	PhoneNumber   string                 `json:"phone_number"`
	PhoneVerified bool                   `json:"phone_verified"`
	UserID        string                 `json:"user_id"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	Identities    []Identity             `json:"identities"`
	AppMetadata   map[string]interface{} `json:"app_metadata"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	Picture       string                 `json:"picture"`
	Name          string                 `json:"name"`
	Nickname      string                 `json:"nickname"`
	Multifactor   []string               `json:"multifactor"`
	LastIP        string                 `json:"last_ip"`
	LastLogin     string                 `json:"last_login"`
	LoginsCount   int                    `json:"logins_count"`
	Blocked       bool                   `json:"blocked"`
	GivenName     string                 `json:"given_name"`
	FamilyName    string                 `json:"family_name"`
}

type Identity struct {
	Connection string `json:"connection,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	Provider   string `json:"provider,omitempty"`
	IsSocial   bool   `json:"isSocial,omitempty"`
}

type GetUserRequestData struct {
	UserID string `json:"user_id"`
}

type GetUserResponseData struct {
	Email         string                 `json:"email"`
	EmailVerified bool                   `json:"email_verified"`
	Username      string                 `json:"username"`
	PhoneNumber   string                 `json:"phone_number"`
	PhoneVerified bool                   `json:"phone_verified"`
	UserID        string                 `json:"user_id"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	Identities    []Identity             `json:"identities"`
	AppMetadata   map[string]interface{} `json:"app_metadata"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	Picture       string                 `json:"picture"`
	Name          string                 `json:"name"`
	Nickname      string                 `json:"nickname"`
	Multifactor   []string               `json:"multifactor"`
	LastIP        string                 `json:"last_ip"`
	LastLogin     string                 `json:"last_login"`
	LoginsCount   int                    `json:"logins_count"`
	Blocked       bool                   `json:"blocked"`
	GivenName     string                 `json:"given_name"`
	FamilyName    string                 `json:"family_name"`
}

type SendVerificationEmailRequestData struct {
	UserID string `json:"user_id"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorTag   string `json:"error"`
	Message    string `json:"message"`
	ErrorCode  string `json:"errorCode"`
}

func (er ErrorResponse) Error() string {
	return fmt.Sprintf("error code %d: %s", er.StatusCode, er.Message)
}
