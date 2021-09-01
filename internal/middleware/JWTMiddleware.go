package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/internal/models"
)

// SetUserData is middleware function, get jwt.Token and set user login into context
func SetUserData(config middleware.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var token interface{}
			token = c.Get("Data")

			userLogin := token.(*jwt.Token).Claims.(*models.CustomClaims).Login

			c.Set("Login", userLogin)

			return next(c)
		}
	}
}
