package authenticationDomainEntity

import "time"

type SignUpRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string     `json:"access_token"`
	ExpiredTime *time.Time `json:"expired_time"`
}
