package repository

import (
	"eulabs/products/internal/produtos"
	"eulabs/products/internal/produtos/entities"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) produtos.Repository {
	return &repository{db: db}
}

func (r *repository) GetById(id int) (temp *entities.Produto, err error) {

	query, err := r.db.Queryx(`
	select
		p.id ,
		p.descricao ,
		pc.id ,
		pc.descricao,
		c.id ,
		c.descricao 
	from
		produto p
	join produto_categoria pc on
		pc.id = p.produto_categoria_id
	join categoria c on
		c.id = pc.produto_categoria_id
	where p.id = ?`, id)

	for query.Next() {
		temp = &entities.Produto{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
			&temp.Categoria.Id,
			&temp.Categoria.Descricao,
			&temp.Categoria.Categoria.Id,
			&temp.Categoria.Categoria.Descricao,
		)
	}

	return
}

func (r *repository) GetProdutoCategoriaById(id int) (temp *entities.ProdutoCategoria, err error) {
	query, err := r.db.Queryx(`
	select
		pc.id ,
		pc.descricao,
		c.id ,
		c.descricao 
	from
		produto_categoria pc 
	join categoria c on
		c.id = pc.produto_categoria_id
	where pc.id = ?`, id)

	for query.Next() {
		temp = &entities.ProdutoCategoria{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
			&temp.Categoria.Id,
			&temp.Categoria.Descricao,
		)
	}

	return

}

func (r *repository) GetCategoriaById(id int) (temp *entities.Categoria, err error) {
	query, err := r.db.Queryx(`
	select
		c.id ,
		c.descricao 
	from
		categoria c 
	where c.id = ?`, id)

	for query.Next() {
		temp = &entities.Categoria{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
		)
	}

	return

}

func (r *repository) SaveProduto(produto entities.Produto) (err error) {
	tx := r.db.MustBegin()
	defer tx.Commit()

	_, err = r.db.NamedExec(`UPDATE produto SET descricao=:descricao WHERE id=:id`, produto)
	if err != nil {
		tx.Rollback()
	}

	return
}

func (r *repository) InsertProduto(produto entities.Produto) (err error) {
	tx := r.db.MustBegin()
	defer tx.Commit()

	_, err = r.db.NamedExec(`INSERT INTO produto(descricao, produto_categoria_id) VALUES(:descricao, :categoria.id)`, produto)
	if err != nil {
		tx.Rollback()
	}

	return
}

func (r *repository) DeleteById(id int) (temp *entities.Produto, err error) {

	temp, err = r.GetById(id)
	if err != nil {
		return
	}

	tx := r.db.MustBegin()
	defer tx.Commit()

	_, err = r.db.Exec(`DELETE FROM produto WHERE id=?`, id)

	if err != nil {
		tx.Rollback()
	}

	return
}

func (r *repository) GetAll() (result []*entities.Produto, err error) {
	query, err := r.db.Queryx(`
	select
		p.id ,
		p.descricao ,
		pc.id ,
		pc.descricao,
		c.id ,
		c.descricao 
	from
		produto p
	join produto_categoria pc on
		pc.id = p.produto_categoria_id
	join categoria c on
		c.id = pc.produto_categoria_id`)

	for query.Next() {
		temp := &entities.Produto{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
			&temp.Categoria.Id,
			&temp.Categoria.Descricao,
			&temp.Categoria.Categoria.Id,
			&temp.Categoria.Categoria.Descricao,
		)
		result = append(result, temp)
	}

	return
}

func (r *repository) GetAllProdutoCategoria() (result []*entities.ProdutoCategoria, err error) {
	query, err := r.db.Queryx(`
	select
		pc.id ,
		pc.descricao,
		c.id ,
		c.descricao 
	from
		produto_categoria pc
	join categoria c on
		c.id = pc.produto_categoria_id`)

	for query.Next() {
		temp := &entities.ProdutoCategoria{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
			&temp.Categoria.Id,
			&temp.Categoria.Descricao,
		)
		result = append(result, temp)
	}

	return
}

func (r *repository) GetAllCategoria() (result []*entities.Categoria, err error) {
	query, err := r.db.Queryx(`
	select
		c.id ,
		c.descricao 
	from
		categoria c`)

	for query.Next() {
		temp := &entities.Categoria{}
		query.Scan(
			&temp.Id,
			&temp.Descricao,
		)
		result = append(result, temp)
	}

	return
}
