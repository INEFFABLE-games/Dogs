package handler

import (
	"context"
	"fmt"
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const userHandler = "user"

// UserHandler creates new user handler.
type UserHandler struct {
	userService *service.UserService
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

	token, err := u.userService.Login(ctx, usr)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": userHandler,
			"action":  "create",
		}).Errorf(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": userHandler,
		"action":  "login",
	}).Debug(fmt.Sprintf("User logined with token: %s", token))

	return c.String(http.StatusOK, fmt.Sprintf("User logined with token: %s", token))
}

// Logout handle logout request from echo.
func (u *UserHandler) Logout(c echo.Context) error {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

	err := u.userService.Logout(c.Request().Context(), token)
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
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
