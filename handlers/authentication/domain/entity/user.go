package authenticationDomainEntity

type User struct {
	Username string
	Password string
	Email    string
}

type UserProfile struct {
	Firstname string
	Lastname  string
	Gender    string
}
