package authenticationRepository

import (
	"context"

	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/ahsansandiah/dealls-test/packages/log"
	"github.com/ahsansandiah/dealls-test/packages/manager"
	"gorm.io/gorm"

	authDomainInterface "github.com/ahsansandiah/dealls-test/handlers/authentication/domain"
	authDomainEntity "github.com/ahsansandiah/dealls-test/handlers/authentication/domain/entity"
)

type Auth struct {
	DB  *gorm.DB
	log log.Log
	cfg *config.Config
}

func NewAuthRepository(mgr manager.Manager) authDomainInterface.AuthRepository {
	repo := new(Auth)
	repo.DB = mgr.GetGorm()
	repo.log = mgr.GetLog()
	repo.cfg = mgr.GetConfig()

	return repo
}

func (repo *Auth) CreateUser(ctx context.Context, data *authDomainEntity.User) (*authDomainEntity.User, error) {
	user := repo.DB.Create(&data)
	if user.Error != nil {
		repo.log.ErrorLog(ctx, user.Error)
		return nil, user.Error
	}

	return nil, nil
}
