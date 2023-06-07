package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsAunthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
	ErrorHandler: func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Sorry, access denied")
	},
})
