package authenticationDomainEntity

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfile struct {
	UserID    int64      `json:"user_id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Gender    string     `json:"gender"`
	Birthdate *time.Time `json:"birthdate"`
	Bio       string     `json:"bio"`
	Location  string     `json:"location"`
	Premium   bool       `json:"premium"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
