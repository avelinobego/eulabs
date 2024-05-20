package repository_test

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/produtos/entities"
	"eulabs/products/internal/produtos/repository"
	"eulabs/products/pkg/database"
	"fmt"
	"testing"
)

var cfg *config.Config
var err error

func init() {
	cfg, err = config.Load("config_local", "../../../")
}

func Test_Repository(t *testing.T) {
	t.Run("teste obter config", func(t *testing.T) {
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("Testando config %v\n", cfg)
	})

	t.Run("testar conexao", func(t *testing.T) {
		db, err := database.NewDatabase(cfg)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("Conn %v\n", db)
	})

	t.Run("testar busca por id", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)
		prod, err := repo.GetById(6)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%v\n", prod)
	})

	t.Run("produto categoria por id", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)
		categoria, err := repo.GetProdutoCategoriaById(7)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%v\n", categoria)
	})

	t.Run("categoria por id", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)
		categoria, err := repo.GetCategoriaById(3)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%v\n", categoria)
	})

	t.Run("alterar produto", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		produto := entities.Produto{Id: 2, Descricao: "Cama King Size Ortobom"}
		err := repo.SaveProduto(produto)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%v\n", err)
	})

	t.Run("inserir produto", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		for i := range 10 {
			produto := entities.Produto{Descricao: fmt.Sprintf("Fogão Superluxo %d", i), Categoria: entities.ProdutoCategoria{Id: 5}}
			err := repo.InsertProduto(produto)
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("%v\n", err)
		}

	})

	t.Run("excluir produto", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		for i := range 10 {
			produto, err := repo.DeleteById(i + 8)
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("%v\n", produto)
		}

	})

	t.Run("todos produtos", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		all, err := repo.GetAll()
		if err != nil {
			t.Error(err)
			return
		}

		for _, p := range all {
			t.Logf("%v\n", p)
		}

	})

	t.Run("todos produtos categoria", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		all, err := repo.GetAllProdutoCategoria()
		if err != nil {
			t.Error(err)
			return
		}

		for _, p := range all {
			t.Logf("%v\n", p)
		}

	})

	t.Run("todos categoria", func(t *testing.T) {
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)

		all, err := repo.GetAllCategoria()
		if err != nil {
			t.Error(err)
			return
		}

		for _, p := range all {
			t.Logf("%v\n", p)
		}

	})

}
