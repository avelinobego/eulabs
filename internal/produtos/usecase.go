package produtos

import (
	"eulabs/products/internal/produtos/dtos"
	"eulabs/products/internal/produtos/entities"
)

type UseCase interface {
	GetAll() (resposta []*dtos.Produto)
	GetById(id int) (result dtos.Produto, err error)
	Insert(dtos.ProdutoRequest) (err error)
	Update(dtos.ProdutoRequest) (err error)
	Delete(int) (*entities.Produto, error)
}
