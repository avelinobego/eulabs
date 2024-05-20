package dtos

type ProdutoRequest struct {
	Id          int    `json:"id" binding:"required"`
	Descricao   string `json:"descricao" binding:"required"`
	Idcategoria int    `json:"categoria" binding:"required"`
}

type Produto struct {
	Id        int              `json:"id" binding:"required"`
	Descricao string           `json:"descricao" binding:"required"`
	Categoria ProdutoCategoria `json:"categoria" binding:"required"`
}

type ProdutoCategoria struct {
	Id        int       `json:"id" binding:"required"`
	Descricao string    `json:"descricao" binding:"required"`
	Categoria Categoria `json:"categoria" binding:"required"`
}

type Categoria struct {
	Id        int    `json:"id" binding:"required"`
	Descricao string `json:"descricao" binding:"required"`
}
