package controllers_test

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/user/controllers"
	"eulabs/products/internal/user/repository"
	"eulabs/products/internal/user/usecase"
	"eulabs/products/pkg/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_Handlers(t *testing.T) {
	t.Run("obter token", func(t *testing.T) {
		userJSON := `{"login":"avelino.bego@gmail.com", "password": "1234567890" }`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)

		cfg, _ := config.Load("config_local", "../../../")
		db, _ := database.NewDatabase(cfg)
		repo := repository.NewRepo(db)
		usecase := usecase.NewUseCase(repo, cfg)
		handlers := controllers.NewHandlers(usecase)
		handlers.Login(context)

		t.Logf("%v\n", rec.Body)

	})
}
