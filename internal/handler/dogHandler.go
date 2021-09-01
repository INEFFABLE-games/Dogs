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

const dogHandler = "dog"

const ctxTime = 5

// DogHandler creates new dog handler.
type DogHandler struct {
	dogService *service.DogService
}

// Create func for echo request.
func (h *DogHandler) Create(c echo.Context) error {
	dog := models.Dog{}

	userLogin := c.Get("Login")

	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

	err := c.Bind(&dog)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "bind model",
		}).Errorf("dogHandler: unable to bind dog data %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(dog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer cancel()

	err = h.dogService.Create(ctx, dog, token, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "create",
		}).Errorf("unable to create dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": dogHandler,
		"action":  "create",
	}).Debug("dog has been created")

	return c.String(http.StatusOK, "Dog was created")
}

// Get func for echo request.
func (h *DogHandler) Get(c echo.Context) error {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

	userLogin := c.Get("Login")

	ctx, cancle := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer cancle()

	resultDog, err := h.dogService.Get(ctx, token, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "get",
		}).Errorf("dogHandler: unable to get dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": dogHandler,
		"action":  "get",
	}).Debugf("replyed dog %v", resultDog)

	return c.JSON(http.StatusOK, resultDog)
}

// Change func for echo request.
func (h *DogHandler) Change(c echo.Context) error {
	dog := models.Dog{}

	userLogin := c.Get("Login")

	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

	err := c.Bind(&dog)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "bind model",
		}).Errorf("unable to bind dog , %v ", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate dog data
	err = c.Validate(dog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancle := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer cancle()

	// check is dog exist
	_, err = h.dogService.Get(ctx, token, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "get",
		}).Errorf("unable to get dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.dogService.Change(ctx, token, dog)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "change",
		}).Errorf("dogHandler: unable to cahnge dog , %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": dogHandler,
		"action":  "change",
	}).Debug("dog has been changed")

	return c.String(http.StatusOK, "Dog has been changed")
}

// Delete func for echo request.
func (h *DogHandler) Delete(c echo.Context) error {

	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

	userLogin := c.Get("Login")

	ctx, canlce := context.WithTimeout(c.Request().Context(), time.Second*ctxTime)
	defer canlce()

	err := h.dogService.Delete(ctx, token, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": dogHandler,
			"action":  "delete",
		}).Errorf("unable to delete dog , %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": dogHandler,
		"action":  "delete",
	}).Debug("dog was deleted")

	return c.String(http.StatusOK, "Dog was deleted")
}

// NewDogHandler create new handler for echo.
func NewDogHandler(dogService *service.DogService) *DogHandler {
	return &DogHandler{dogService: dogService}
}
