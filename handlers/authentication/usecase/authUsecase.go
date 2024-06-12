package authenticationUsecase

import (
	"context"
	"fmt"

	jwtAuth "github.com/ahsansandiah/dealls-test/packages/auth/jwt"
	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/ahsansandiah/dealls-test/packages/log"
	"github.com/ahsansandiah/dealls-test/packages/manager"

	authDomainInterface "github.com/ahsansandiah/dealls-test/handlers/authentication/domain"
	authDomainEntity "github.com/ahsansandiah/dealls-test/handlers/authentication/domain/entity"
)

type Authentication struct {
	jwt jwtAuth.Jwt
	log log.Log
	cfg *config.Config
}

func NewAuthUsecase(mgr manager.Manager) authDomainInterface.AuthUsecase {
	usecase := new(Authentication)
	usecase.jwt = mgr.GetJwt()
	usecase.log = mgr.GetLog()
	usecase.cfg = mgr.GetConfig()

	return usecase
}

func (uc *Authentication) SignUp(ctx context.Context, data *authDomainEntity.SignUpRequest) error {

	fmt.Println(data)
	return nil
}
