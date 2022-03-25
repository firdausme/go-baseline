package controller

import (
	"github.com/labstack/echo/v4"
	"go-baseline/model"
	"go-baseline/service"
	"net/http"
)

type LoginController struct {
	UserService service.UserService
}

func NewLoginController(userService *service.UserService) LoginController {
	return LoginController{UserService: *userService}
}

func (controller *LoginController) Route(route *echo.Group) {
	route.POST("/login", controller.Login)
}

func (controller *LoginController) Login(c echo.Context) error {

	request := model.LoginRequest{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	response := controller.UserService.Login(request)

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	})
}
