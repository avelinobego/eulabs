package usecase_test

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/user/dtos"
	"eulabs/products/internal/user/repository"
	"eulabs/products/internal/user/usecase"
	"eulabs/products/pkg/database"
	"testing"
)

func Test_Usecase(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		cfg, _ := config.Load("config_local")
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)
		u := usecase.NewUseCase(repo, cfg)

		request := &dtos.UserLoginRequest{Login: "avelino.bego@gmail.com", Password: "1234567890"}
		token, err := u.Login(request)
		if err != nil {
			t.Error(err)
		}
		t.Log(token)

	})
}
