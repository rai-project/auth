package api

import "fmt"

// User ...
type User struct {
	Email         string                 `json:"email,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	Username      string                 `json:"username,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	UserID        string                 `json:"user_id,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	UpdatedAt     string                 `json:"updated_at,omitempty"`
	Identities    []Identity             `json:"identities,omitempty"`
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	Picture       string                 `json:"picture,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Nickname      string                 `json:"nickname,omitempty"`
	Multifactor   []string               `json:"multifactor,omitempty"`
	LastIP        string                 `json:"last_ip,omitempty"`
	LastLogin     string                 `json:"last_login,omitempty"`
	LoginsCount   int                    `json:"logins_count,omitempty"`
	Blocked       bool                   `json:"blocked,omitempty"`
	GivenName     string                 `json:"given_name,omitempty"`
	FamilyName    string                 `json:"family_name,omitempty"`
}

// Identity ...
type Identity struct {
	Connection string `json:"connection,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	Provider   string `json:"provider,omitempty"`
	IsSocial   bool   `json:"isSocial,omitempty"`
}

// UserPage ...
type UserPage struct {
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
	Length int    `json:"length"`
	Users  []User `json:"users"`
}

// CreateUserRequestData ...
type CreateUserRequestData struct {
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
	Connection    string                 `json:"connection"`
	Email         string                 `json:"email,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	Password      string                 `json:"password,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	Username      string                 `json:"username,omitempty"`
	VerifyEmail   bool                   `json:"verify_email,omitempty"`
	GivenName     string                 `json:"-"`
}

// UpdateUserRequestData ...
type UpdateUserRequestData struct {
	AppMetadata       map[string]interface{} `json:"app_metadata,omitempty"`
	Blocked           bool                   `json:"blocked,omitempty"`
	ClientID          string                 `json:"client_id,omitempty"`
	Connection        string                 `json:"connection,omitempty"`
	Email             string                 `json:"email,omitempty"`
	EmailVerified     bool                   `json:"email_verified,omitempty"`
	Password          string                 `json:"password,omitempty"`
	PhoneNumber       string                 `json:"phone_number,omitempty"`
	PhoneVerified     bool                   `json:"phone_verified,omitempty"`
	UserMetadata      map[string]interface{} `json:"user_metadata,omitempty"`
	Username          string                 `json:"username,omitempty"`
	VerifyEmail       bool                   `json:"verify_email,omitempty"`
	VerifyPassword    bool                   `json:"verify_password,omitempty"`
	VerifyPhoneNumber bool                   `json:"verify_phone_number,omitempty"`
	GivenName         string                 `json:"-"`
}

// GetUserRequestData ...
type GetUserRequestData struct {
	UserID string `json:"user_id"`
}

// SendVerificationEmailRequestData ...
type SendVerificationEmailRequestData struct {
	UserID string `json:"user_id"`
}

// ErrorResponse ...
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorTag   string `json:"error"`
	Message    string `json:"message"`
	ErrorCode  string `json:"errorCode"`
}

// Error ...
func (er ErrorResponse) Error() string {
	return fmt.Sprintf("error code %d: %s", er.StatusCode, er.Message)
}
