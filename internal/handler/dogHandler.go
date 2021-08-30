package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"time"
)

// DogHandler creates new dog handler.
type DogHandler struct {
	dogService *service.DogService
}

const ctxtime = 5

// Create func for echo request.
func (h *DogHandler) Create(c echo.Context) error {
	log.SetLevel(log.DebugLevel)
	dog := models.NewDog("", "")

	err := c.Bind(&dog)
	if err != nil {
		log.Errorf("dogHandler: unable to bind dog data %v,", err)
	}

	err = c.Validate(dog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*ctxtime)
	defer cancel()

	err = h.dogService.Create(ctx, dog)
	if err != nil {
		log.Errorf("unable to create dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": "dog",
		"action":  "create",
	}).Debug("dog has been created")

	return echo.NewHTTPError(http.StatusOK, "Dog was created")
}

// Get func for echo request.
func (h *DogHandler) Get(c echo.Context) error {
	name := c.QueryParam("name")

	ctx, cancle := context.WithTimeout(c.Request().Context(), time.Second*ctxtime)
	defer cancle()

	resultDog, err := h.dogService.Get(ctx, name)
	if err != nil {
		log.Errorf("dogHandler: unable to get dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": "dog",
		"action":  "get",
	}).Debugf("replyed dog %v", resultDog)

	return echo.NewHTTPError(http.StatusOK, resultDog)
}

// Change func for echo request.
func (h *DogHandler) Change(c echo.Context) error {
	dog := models.NewDog("", "")

	err := c.Bind(&dog)
	if err != nil {
		log.Errorf("unable to bind dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	name := c.QueryParam("name")

	// validate dog data
	err = c.Validate(dog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancle := context.WithTimeout(c.Request().Context(), time.Second*ctxtime)
	defer cancle()

	// check is dog exist
	_, err = h.dogService.Get(ctx, name)
	if err != nil {
		log.Errorf("unable to get dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.dogService.Change(ctx, name, dog)
	if err != nil {
		log.Errorf("dogHandler: unable to cahnge dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": "dog",
		"action":  "cahnge",
	}).Debug("dog has been changed")
	return echo.NewHTTPError(http.StatusOK, "Dog has been changed")
}

// Delete func for echo request.
func (h *DogHandler) Delete(c echo.Context) error {
	name := c.QueryParam("name")

	ctx, canlce := context.WithTimeout(c.Request().Context(), time.Second*ctxtime)
	defer canlce()

	// check is dog exist
	_, err := h.dogService.Get(ctx, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Dog doesn't exist")
	}

	err = h.dogService.Delete(ctx, name)
	if err != nil {
		log.Errorf("unable to delete dog %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.WithFields(log.Fields{
		"handler": "dog",
		"action":  "delete",
	}).Debug("dog was deleted")
	return echo.NewHTTPError(http.StatusOK, "Dog was deleted")
}

// NewDogHandler create new handler for echo.
func NewDogHandler(dogService *service.DogService) DogHandler {
	return DogHandler{dogService: dogService}
}
