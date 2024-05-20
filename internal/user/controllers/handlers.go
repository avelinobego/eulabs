package controllers

import (
	"eulabs/products/internal/user"
	"eulabs/products/internal/user/dtos"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlers struct {
	usecase user.UseCase
}

func NewHandlers(u user.UseCase) *handlers {
	return &handlers{usecase: u}
}

func (h *handlers) SetHandlers(domain *echo.Group) {
	domain.GET("/login", h.Login)
}

func (h *handlers) Login(c echo.Context) (err error) {

	user := &dtos.UserLoginRequest{}
	c.Bind(user)
	reponse, err := h.usecase.Login(user)
	if err != nil {
		return
	}

	err = c.JSON(http.StatusOK, reponse)

	return
}
