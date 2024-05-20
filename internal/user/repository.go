package user

import "eulabs/products/internal/user/dtos"

type Repository interface {
	GetUserByLogin(login string) (response *dtos.UserLoginRequest, err error)
}
