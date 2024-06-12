package authenticationUsecase

import (
	"context"

	jwtAuth "github.com/ahsansandiah/dealls-test/packages/auth/jwt"
	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/ahsansandiah/dealls-test/packages/log"
	"github.com/ahsansandiah/dealls-test/packages/manager"
	"golang.org/x/crypto/bcrypt"

	authDomainInterface "github.com/ahsansandiah/dealls-test/handlers/authentication/domain"
	authDomainEntity "github.com/ahsansandiah/dealls-test/handlers/authentication/domain/entity"
	authRepository "github.com/ahsansandiah/dealls-test/handlers/authentication/repository"
)

type Authentication struct {
	jwt  jwtAuth.Jwt
	log  log.Log
	cfg  *config.Config
	repo authDomainInterface.AuthRepository
}

func NewAuthUsecase(mgr manager.Manager) authDomainInterface.AuthUsecase {
	usecase := new(Authentication)
	usecase.jwt = mgr.GetJwt()
	usecase.log = mgr.GetLog()
	usecase.cfg = mgr.GetConfig()
	usecase.repo = authRepository.NewAuthRepository(mgr)

	return usecase
}

func (uc *Authentication) SignUp(ctx context.Context, data *authDomainEntity.SignUpRequest) (*authDomainEntity.UserProfile, error) {
	password := data.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, err
	}

	userData := authDomainEntity.User{
		Username: data.Username,
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	user, err := uc.repo.CreateUser(ctx, &userData)
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, err
	}

	userProfileData := authDomainEntity.UserProfile{
		UserID:    user.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Gender:    data.Gender,
		Premium:   false,
	}

	err = uc.repo.CreateUserProfile(ctx, &userProfileData)
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, err
	}

	return &userProfileData, nil
}

func (uc *Authentication) Login(ctx context.Context, data *authDomainEntity.LoginRequest) (*authDomainEntity.LoginResponse, error) {
	// check user exists
	user, err := uc.repo.GetUserByUsername(ctx, data.Username)
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, err
	}

	// check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, nil
	}

	// generate JWT Token
	dataJwt := &jwtAuth.JwtData{
		UserID: user.ID,
	}

	accessToken, expiredTime, err := uc.jwt.GenerateToken(dataJwt, false)
	if err != nil {
		uc.log.ErrorLog(ctx, err)
		return nil, err
	}

	results := &authDomainEntity.LoginResponse{
		AccessToken: accessToken,
		ExpiredTime: expiredTime,
	}

	return results, nil
}
