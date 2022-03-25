package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
	"go-baseline/config"
	Env "go-baseline/config/env"
	"go-baseline/constant"
	"go-baseline/middleware"
)

var e *echo.Echo

func main() {

	// load log4go configuration file
	log.LoadConfiguration("./log4go.json")

	e = echo.New()

	// setup database
	driver := config.SetupDB()

	// add middleware log
	e.Use(middleware.Log)
	// e.HTTPErrorHandler = middleware.Error

	env := Env.Get()

	// setup repositories, services, controllers
	config.NewServices(driver, e)

	e.Logger.Fatal(e.Start(env[constant.AppServerPort]))
}
