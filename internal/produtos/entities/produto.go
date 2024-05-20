package entities

type Produto struct {
	Id        int
	Descricao string
	Categoria ProdutoCategoria
}

type ProdutoCategoria struct {
	Id        int
	Descricao string
	Categoria Categoria
}

type Categoria struct {
	Id        int
	Descricao string
}
