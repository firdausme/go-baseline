package controller

import (
	"github.com/labstack/echo/v4"
	"go-baseline/middleware"
	"go-baseline/model"
	"go-baseline/service"
	"go-baseline/utils"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(route *echo.Group) {
	route.GET("/user", controller.FindAll, middleware.IsAuthenticated)
	route.POST("/user", controller.Create)
	route.PUT("/user", controller.Update, middleware.IsAuthenticated)
	route.DELETE("/user", controller.Delete, middleware.IsAuthenticated)
}

func (controller *UserController) FindAll(c echo.Context) error {
	data := controller.UserService.FindAll()
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func (controller *UserController) Create(c echo.Context) error {

	request := model.CreateUserRequest{
		Username: c.FormValue("username"),
		Password: utils.HashPassword(c.FormValue("password")),
	}

	response := controller.UserService.Create(request)
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

func (controller *UserController) Update(c echo.Context) error {
	request := model.UpdateUserRequest{
		Id:       c.FormValue("id"),
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	response := controller.UserService.Update(request)
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

func (controller *UserController) Delete(c echo.Context) error {
	request := model.DeleteUserRequest{
		Id: c.QueryParam("id"),
	}

	controller.UserService.Delete(request)

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    "",
	})
}
