package controllers

import (
	"eulabs/products/internal/produtos"
	"eulabs/products/internal/produtos/dtos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlers struct {
	usecase produtos.UseCase
}

func NewHandlers(u produtos.UseCase) *handlers {
	return &handlers{usecase: u}
}

func (h *handlers) SetHandlers(domain *echo.Group) {
	domain.GET("/all", h.GetAll)
	domain.GET("/:id", h.GetById)
	domain.POST("/inserir", h.Insert)
	domain.PUT("/atualizar", h.Update)
	domain.DELETE("/delete/:id", h.Delete)
}

func (h *handlers) GetAll(c echo.Context) (err error) {

	all := h.usecase.GetAll()
	reposta := make([]dtos.Produto, 0, len(all))

	for _, e := range all {
		reposta = append(reposta, dtos.Produto{
			Id:        e.Id,
			Descricao: e.Descricao,
			Categoria: dtos.ProdutoCategoria{
				Id:        e.Categoria.Id,
				Descricao: e.Categoria.Descricao,
				Categoria: dtos.Categoria{
					Id:        e.Categoria.Categoria.Id,
					Descricao: e.Categoria.Categoria.Descricao,
				},
			},
		})
	}

	err = c.JSON(http.StatusOK, reposta)
	return
}

func (h *handlers) GetById(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	produto, err := h.usecase.GetById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, produto)
	return
}

func (h *handlers) Insert(c echo.Context) (err error) {
	dto := dtos.ProdutoRequest{}
	c.Bind(&dto)
	err = h.usecase.Insert(dto)
	return
}

func (h *handlers) Update(c echo.Context) (err error) {
	dto := dtos.ProdutoRequest{}
	c.Bind(&dto)
	err = h.usecase.Update(dto)
	return
}

func (h *handlers) Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	apagado, err := h.usecase.Delete(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, apagado)
	return
}
