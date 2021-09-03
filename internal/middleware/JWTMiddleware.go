package middleware

import (
	"Dogs/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

// AuthenticateToken is middleware function, get jwt.Token and set user login into context
func AuthenticateToken(config middleware.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, err := jwt.ParseWithClaims(
				strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "),
				&models.CustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return config.SigningKey, nil
				},
			)

			if err != nil {
				return echo.NewHTTPError(
					http.StatusBadRequest,
					"unable to parse token",
				)
			}

			claim, ok := token.Claims.(*models.CustomClaims)
			if !ok {
				return echo.NewHTTPError(
					http.StatusBadRequest,
					"couldn't get claim from token",
				)
			}

			c.Set("Login", claim.Login)

			return next(c)
		}
	}
}
