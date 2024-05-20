package usercase

import (
	"eulabs/products/internal/produtos"
	"eulabs/products/internal/produtos/dtos"
	"eulabs/products/internal/produtos/entities"
	"eulabs/products/pkg/utils"

	config "eulabs/products/etc"
)

type usecase struct {
	repository produtos.Repository
	cfg        *config.Config
}

func NewUseCase(repository produtos.Repository, cfg *config.Config) produtos.UseCase {
	return &usecase{repository: repository, cfg: cfg}
}

func (uc *usecase) GetAll() (resposta []*dtos.Produto) {
	all, err := uc.repository.GetAll()
	if utils.Error(err) {
		return nil
	}
	for _, e := range all {
		resposta = append(resposta, &dtos.Produto{
			Id:        e.Id,
			Descricao: e.Descricao,
			Categoria: dtos.ProdutoCategoria{
				Id:        e.Categoria.Id,
				Descricao: e.Categoria.Categoria.Descricao,
				Categoria: dtos.Categoria{
					Id:        e.Categoria.Categoria.Id,
					Descricao: e.Categoria.Descricao,
				},
			}})
	}
	return
}

func (uc *usecase) GetById(id int) (result dtos.Produto, err error) {
	e, err := uc.repository.GetById(id)
	if err != nil {
		return
	}

	result = dtos.Produto{
		Id:        e.Id,
		Descricao: e.Descricao,
		Categoria: dtos.ProdutoCategoria{
			Id:        e.Categoria.Id,
			Descricao: e.Categoria.Categoria.Descricao,
			Categoria: dtos.Categoria{
				Id:        e.Categoria.Categoria.Id,
				Descricao: e.Categoria.Descricao,
			},
		}}

	return
}

func (uc *usecase) Insert(request dtos.ProdutoRequest) (err error) {

	produto_categira, err := uc.repository.GetProdutoCategoriaById(request.Idcategoria)
	if err != nil {
		return
	}
	err = uc.repository.InsertProduto(entities.Produto{
		Descricao: request.Descricao,
		Categoria: *produto_categira,
	})
	return
}

func (uc *usecase) Update(request dtos.ProdutoRequest) (err error) {
	prod, err := uc.GetById(request.Id)
	if err != nil {
		return
	}
	produto_categoria, err := uc.repository.GetProdutoCategoriaById(request.Idcategoria)
	if err != nil {
		return
	}

	uc.repository.SaveProduto(entities.Produto{
		Id:        prod.Id,
		Descricao: request.Descricao,
		Categoria: *produto_categoria,
	})
	return
}

func (uc *usecase) Delete(id int) (apagado *entities.Produto, err error) {
	apagado, err = uc.repository.DeleteById(id)
	return
}
