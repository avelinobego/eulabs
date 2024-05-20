package repository_test

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/user/dtos"
	"eulabs/products/internal/user/repository"
	"eulabs/products/pkg/database"
	"testing"
)

var cfg *config.Config
var err error

func init() {
	cfg, err = config.Load("config_local")
}

func Test_Geral(t *testing.T) {

	t.Run("Login correto", func(t *testing.T) {
		db, err := database.NewDatabase(cfg)
		if err != nil {
			t.Error(err)
		}
		repo := repository.NewRepo(db)
		request := dtos.UserLoginRequest{Login: "avelino.bego@gmail.com", Password: "1234567890"}
		err = request.Validate()
		if err != nil {
			t.Error(err)
		}
		response, err := repo.GetUserByLogin(request.Login)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v", response)
	})
}
