package authenticationDomainInterface

import (
	"context"
	"net/http"

	authDomainEntity "github.com/ahsansandiah/dealls-test/handlers/authentication/domain/entity"
)

type AuthHandler interface {
	SignUp() http.Handler
}

type AuthUsecase interface {
	SignUp(ctx context.Context, data *authDomainEntity.SignUpRequest) error
}

type AuthRepository interface {
	CreateUser(ctx context.Context, data *authDomainEntity.User) (*authDomainEntity.User, error)
}
