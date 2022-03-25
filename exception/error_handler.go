package exception

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandler(c echo.Context, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "BAD_REQUEST", "DATA": err.Error()})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{"status": "INTERNAL_SERVER_ERROR", "DATA": err.Error()})

}
