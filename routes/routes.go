package routes

import (
	"go-auth-jwt/db"
	"go-auth-jwt/handlers"
	"go-auth-jwt/repository"
	"go-auth-jwt/services"
	"github.com/labstack/echo/v4"
	"go-auth-jwt/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	biodataRepository := repository.NewRepository(db.Init())
	biodataService := services.NewService(biodataRepository)
	biodataHandler := handlers.NewBiodataHandler(biodataService)

	loginRepository := repository.NewRepositoryLogin(db.Init())
	loginService := services.NewServiceLogin(loginRepository)
	loginHandler := handlers.NewLoginHandler(loginService)

	e.GET("/biodata", biodataHandler.GetAll, middleware.IsAunthenticated)
	e.GET("/generate-hash/:password", handlers.GenerateHashPassword)
	e.POST("/login", loginHandler.CheckLogin)
	return e
}
