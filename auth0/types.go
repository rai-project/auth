package auth0

import "fmt"

type CreateUserRequestData struct {
	Connection    string                 `json:"connection"`
	Email         string                 `json:"email"`
	Username      string                 `json:"username"`
	Password      string                 `json:"password"`
	PhoneNumber   string                 `json:"phone_number"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	EmailVerified bool                   `json:"email_verified"`
	VerifyEmail   bool                   `json:"verify_email"`
	PhoneVerified bool                   `json:"phone_verified"`
	AppMetadata   map[string]interface{} `json:"app_metadata"`
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
	Connection string `json:"connection"`
	UserID     string `json:"user_id"`
	Provider   string `json:"provider"`
	IsSocial   bool   `json:"isSocial"`
}

type GetUserRequestData struct {
	ID string `json:"user_id"`
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
