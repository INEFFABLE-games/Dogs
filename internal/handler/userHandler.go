package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"time"
)

const userHandler = "user"

// UserHandler creates new user handler.
type UserHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

// Create handle create request from echo.
func (u *UserHandler) Create(c echo.Context) error {
	usr := models.User{}

	err := c.Bind(&usr)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "bind model",
		}).Errorf("userHandler: unable to bind user data %v,", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer cancel()

	err = u.userService.Create(ctx, usr)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "create",
		}).Errorf(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "User created")
}

// Login handle login request from echo.
func (u *UserHandler) Login(c echo.Context) error {
	usr := models.User{}

	err := c.Bind(&usr)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "bind model",
		}).Errorf("userHandler: unable to bind user data %v,", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer cancel()

	err = u.userService.Login(ctx, usr)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "login",
		}).Errorf(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	JWT, RT, err := u.authService.CreateTokens(ctx, usr.Login)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "create tokens",
		}).Errorf(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": userHandler,
		"action":  "login",
	}).Debug(fmt.Sprintf("User logined with tokens: JWT[%s] RT[%s]", JWT, RT))

	return c.String(http.StatusOK, fmt.Sprintf("User logined with tokens: JWT[%s]	RT[%s]", JWT, RT))
}

// Logout handle logout request from echo.
func (u *UserHandler) Logout(c echo.Context) error {

	err := u.userService.Logout(c.Request().Context())
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "logout",
		}).Errorf("unable to logout user %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "User logout success")
}

// NewUserHandler create new handler for echo.
func NewUserHandler(userService *service.UserService, authService *service.AuthService) *UserHandler {
	return &UserHandler{userService: userService, authService: authService}
}
