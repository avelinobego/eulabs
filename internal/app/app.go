package app

import (
	config "eulabs/products/etc"
	produtos_controllers "eulabs/products/internal/produtos/controllers"
	produtos_repo "eulabs/products/internal/produtos/repository"
	produtos_usecase "eulabs/products/internal/produtos/usecase"
	user_controllers "eulabs/products/internal/user/controllers"
	user_repo "eulabs/products/internal/user/repository"
	user_usecase "eulabs/products/internal/user/usecase"
	"eulabs/products/pkg/database"
	"fmt"
	"net/http"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type App struct {
	server *http.Server
}

func NewApp(cfg *config.Config) *App {

	_, err := time.LoadLocation(cfg.Server.TimeZone)
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	// Server e rotas
	e := echo.New()

	grupo_permitido := e.Group("/api/v1")

	// Login
	user_handlers := user_controllers.NewHandlers(
		user_usecase.NewUseCase(
			user_repo.NewRepo(db), cfg))

	user_handlers.SetHandlers(grupo_permitido)

	// Produtos
	produtos_handlers := produtos_controllers.NewHandlers(
		produtos_usecase.NewUseCase(
			produtos_repo.NewRepo(db), cfg))

	grupo_restrito := e.Group("/api/v1/produtos")
	grupo_restrito.Use(echojwt.JWT([]byte(cfg.Authentication.Key)))
	produtos_handlers.SetHandlers(grupo_restrito)
	// Server e rotas

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:     e,
		ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	return &App{server: s}
}

func (a *App) Run() (err error) {
	err = a.server.ListenAndServe()
	return
}
