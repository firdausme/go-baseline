package middleware

import (
	"github.com/labstack/echo/v4/middleware"
	Env "go-baseline/config/env"
	"go-baseline/constant"
)

var jwtSecret = func() string {
	env := Env.Get()
	return env[constant.JwtSecret]
}

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(jwtSecret()),
})
