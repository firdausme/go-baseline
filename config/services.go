package config

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	Env "go-baseline/config/env"
	"go-baseline/constant"
	"go-baseline/controller"
	"go-baseline/repository"
	"go-baseline/service"
)

func NewServices(db *sql.DB, e *echo.Echo) {

	env := Env.Get()
	route := e.Group(env[constant.AppBasePath])

	// user
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)
	userController.Route(route)

	// login
	loginController := controller.NewLoginController(&userService)
	loginController.Route(route)

}
