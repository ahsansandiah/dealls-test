package authenticationDomainEntity

type SignUpRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
