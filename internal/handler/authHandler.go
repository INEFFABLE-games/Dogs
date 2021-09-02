package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strings"
)

// AuthHandler struct for auth handler
type AuthHandler struct {
	authService *service.AuthService
}

// RefreshTokens handler func for refresh request
func (a *AuthHandler) RefreshTokens(c echo.Context) error {

	token, err := jwt.ParseWithClaims(
		strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "),
		&models.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("RefTokKey"), nil
		},
	)

	if err != nil {
		log.WithFields(log.Fields{
			"handler": "auth",
			"action":  "pars token",
		})

		return echo.NewHTTPError(
			http.StatusBadRequest,
			"unable to pars token",
		)
	}

	claim, ok := token.Claims.(*models.CustomClaims)
	if !ok {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"couldn't get claim from token",
		)
	}

	jwtTok, rt, err := a.authService.RefreshTokens(c.Request().Context(), claim.Login)
	if err != nil {

		log.WithFields(log.Fields{
			"handler": "auth",
			"action":  "refresh token",
		})

		return echo.NewHTTPError(
			http.StatusBadRequest,
			"unable to save token in database",
		)
	}

	return c.String(http.StatusOK, fmt.Sprintf("Your new tokens: RT[%s]	JWT[%s]", rt, jwtTok))
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
