package user

import "eulabs/products/internal/user/dtos"

type UseCase interface {
	Login(request *dtos.UserLoginRequest) (response dtos.UserLoginResponse, err error)
}
