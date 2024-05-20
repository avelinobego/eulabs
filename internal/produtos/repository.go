package produtos

import "eulabs/products/internal/produtos/entities"

type Repository interface {
	GetById(id int) (result *entities.Produto, err error)
	GetProdutoCategoriaById(id int) (result *entities.ProdutoCategoria, err error)
	GetCategoriaById(id int) (result *entities.Categoria, err error)
	SaveProduto(produto entities.Produto) (err error)
	InsertProduto(produto entities.Produto) (err error)
	DeleteById(id int) (result *entities.Produto, err error)

	GetAll() (result []*entities.Produto, err error)
	GetAllProdutoCategoria() (result []*entities.ProdutoCategoria, err error)
	GetAllCategoria() (result []*entities.Categoria, err error)
}
