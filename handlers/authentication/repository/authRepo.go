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
	user := authDomainEntity.User{}
	result := repo.DB.Create(&data).Scan(&user)
	if result.Error != nil {
		repo.log.ErrorLog(ctx, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (repo *Auth) CreateUserProfile(ctx context.Context, userProfile *authDomainEntity.UserProfile) error {
	err := repo.DB.Create(&userProfile).Error
	if err != nil {
		repo.log.ErrorLog(ctx, err)
		return err
	}

	return err
}

func (repo *Auth) GetUserByUsername(ctx context.Context, username string) (*authDomainEntity.User, error) {
	user := &authDomainEntity.User{}
	var total int64

	res := repo.DB.Where("username = ?", username).Find(user)

	err := res.Error
	if err != nil {
		repo.log.ErrorLog(ctx, err)
		return nil, err
	}

	res.Count(&total)
	if total == 0 {
		repo.log.ErrorLog(ctx, err)
		return nil, err
	}

	return user, nil
}
