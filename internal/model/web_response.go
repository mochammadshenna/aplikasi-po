package model

type GoogleCredential struct {
	Credential string `json:"credential"`
}

// You might already have AuthAdminResponse, but make sure it includes these fields
type AuthAdminResponse struct {
	Token string       `json:"access_token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	Id      int64  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
