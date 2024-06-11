package jwtAuth

import (
	"fmt"
	"strings"
	"time"

	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/golang-jwt/jwt"
)

type Jwt interface {
	GenerateToken(data *JwtData, isVA bool) (string, *time.Time, error)
	ExtractJwtToken(token string) (*JwtData, error)
	VerifyAccessToken(token string, combinedKey string) (*JwtData, error)
}

type Options struct {
	secretKey             string
	accessTokenDuration   int
	accessTokenVADuration int
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.secretKey = cfg.JwtSecretKey
	opt.accessTokenDuration = cfg.JwtAccessTokenDuration
	opt.accessTokenVADuration = cfg.JwtAccessTokenVADuration

	return opt
}

func (o *Options) GenerateToken(data *JwtData, isVA bool) (string, *time.Time, error) {
	jwtPayload := &JwtPayload{
		Token:  data.Token,
		UserID: data.UserID,
	}

	expiredTime := time.Now().Local().Add(time.Second * time.Duration(o.accessTokenDuration))
	if isVA {
		expiredTime = data.ExpiredAt
	}

	jwtPayload.StandardClaims.ExpiresAt = expiredTime.Unix()
	jwtPayload.StandardClaims.NotBefore = jwt.TimeFunc().Local().Unix()
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtPayload)

	combineKey := o.secretKey
	accessToken, err := acToken.SignedString([]byte(combineKey))
	if err != nil {
		return "", nil, err
	}

	return accessToken, &expiredTime, nil
}

func (o *Options) ExtractJwtToken(token string) (*JwtData, error) {
	parsedToken, err := jwt.Parse(token, nil)
	if strings.Contains(err.Error(), "invalid number of segments") {
		return nil, err
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	jwtData := &JwtData{
		Token: claims["token"].(string),
	}

	return jwtData, nil
}

func (o *Options) VerifyAccessToken(token string, combinedKey string) (*JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(combinedKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	jwtData := &JwtData{
		Token: claims["token"].(string),
	}

	return jwtData, nil
}
