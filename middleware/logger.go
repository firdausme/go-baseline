package middleware

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type Request struct {
	Method      string            `json:"method"`
	Uri         string            `json:"uri"`
	Ip          string            `json:"ip"`
	ContentType string            `json:"content-type"`
	Body        map[string]string `json:"body"`
}

func getContentType(c echo.Context) string {
	return strings.Split(c.Request().Header.Get("content-type"), ";")[0]
}

func getBody(c echo.Context) map[string]string {

	params := make(map[string]string)

	ctype := getContentType(c)

	if strings.Contains(ctype, "application/x-www-form-urlencoded") {

		err := c.Request().ParseForm()
		if err != nil {
			return params
		}
		for k := range c.Request().Form {
			params[k] = c.FormValue(k)
		}

	} else if strings.Contains(ctype, "multipart/form-data") {

		err := c.Request().ParseMultipartForm(math.MaxInt64)
		if err != nil {
			return params
		}
		for k := range c.Request().Form {
			params[k] = c.FormValue(k)
		}

	} else if strings.Contains(ctype, "application/json") {

		body := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&body)

		if err != nil {
			return params
		}

		for k, v := range body {

			val := reflect.ValueOf(v)

			switch val.Kind() {
			case reflect.Map:
				result, _ := json.Marshal(val.Interface())
				params[k] = string(result)
			case reflect.Slice:
				result, _ := json.Marshal(val.Interface())
				params[k] = string(result)
			case reflect.Float64:
				params[k] = strconv.FormatFloat(val.Float(), 'f', -1, 64)
			case reflect.Bool:
				params[k] = strconv.FormatBool(val.Bool())
			default:
				params[k] = val.String()
			}
		}
	}

	return params
}

func logEntry(c echo.Context) string {

	req := Request{
		c.Request().Method,
		c.Request().URL.String(),
		c.RealIP(),
		getContentType(c),
		getBody(c),
	}

	result, _ := json.Marshal(req)

	return string(result)
}

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		_, _ = c.FormParams()
		log.Info(logEntry(c))

		return next(c)
	}
}
