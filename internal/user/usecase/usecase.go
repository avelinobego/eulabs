package usecase

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/user"
	"eulabs/products/internal/user/dtos"
	"eulabs/products/pkg/jwt"
	"fmt"
)

type usecase struct {
	repository user.Repository
	cfg        *config.Config
}

func NewUseCase(repository user.Repository, cfg *config.Config) user.UseCase {
	return &usecase{repository: repository, cfg: cfg}
}

func (u *usecase) Login(request *dtos.UserLoginRequest) (response dtos.UserLoginResponse, err error) {
	err = request.Validate()
	if err != nil {
		return
	}

	user, err := u.repository.GetUserByLogin(request.Login)
	if err != nil {
		return
	}

	if request.Login != user.Login || request.Password != user.Password {
		err = fmt.Errorf("ivalid username or password")
		return
	}

	response, err = jwt.Generate(u.cfg.Authentication.Key, request.Login)
	if err != nil {
		return
	}

	return
}
