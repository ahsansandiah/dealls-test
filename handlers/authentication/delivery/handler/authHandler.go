package authenticationHandler

import (
	"encoding/json"
	"net/http"

	res "github.com/ahsansandiah/dealls-test/packages/json"
	"github.com/ahsansandiah/dealls-test/packages/log"
	"github.com/ahsansandiah/dealls-test/packages/manager"

	authDomainInterface "github.com/ahsansandiah/dealls-test/handlers/authentication/domain"
	authDomainEntity "github.com/ahsansandiah/dealls-test/handlers/authentication/domain/entity"
	authenticationUsecase "github.com/ahsansandiah/dealls-test/handlers/authentication/usecase"
)

type Auth struct {
	Usecase authDomainInterface.AuthUsecase
	log     log.Log
	Json    res.Json
}

func NewAuthHandler(mgr manager.Manager) authDomainInterface.AuthHandler {
	handler := new(Auth)
	handler.Usecase = authenticationUsecase.NewAuthUsecase(mgr)
	handler.Json = mgr.GetJson()

	return handler
}

func (h *Auth) SignUp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req *authDomainEntity.SignUpRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.Json.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		_, err := h.Usecase.SignUp(ctx, req)
		if err != nil {
			h.Json.ErrorResponse(w, r, http.StatusInternalServerError, err)
			return
		}

		h.Json.SuccessResponse(w, r, http.StatusCreated, "success sign up", "")
	})
}

func (h *Auth) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req *authDomainEntity.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.Json.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		results, err := h.Usecase.Login(ctx, req)
		if err != nil {
			h.Json.ErrorResponse(w, r, http.StatusInternalServerError, err)
			return
		}

		h.Json.SuccessResponse(w, r, http.StatusCreated, "success login", &results)
	})
}
