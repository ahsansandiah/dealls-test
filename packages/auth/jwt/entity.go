package jwtAuth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtData struct {
	Token     string    `json:"token"`
	UserID    int64     `json:"user_id"`
	ExpiredAt time.Time `json:"expired_at"`
}

type JwtPayload struct {
	Token  string `json:"token"`
	UserID int64  `json:"ui"`
	jwt.StandardClaims
}
