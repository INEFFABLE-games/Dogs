package handler

import (
	"Dogs/internal/models"
	"Dogs/internal/service"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const birdHandler = "bird"

// BirdHandler creates new dog handler.
type BirdHandler struct {
	birdService *service.BirdService
}

// Create func for echo request.
func (b *BirdHandler) Create(c echo.Context) error {

	userLogin := c.Get("Login")

	bird := models.Bird{}

	err := c.Bind(&bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "bind model",
		}).Errorf("dogHandler: unable to bind dog data %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(bird)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bird.Owner = fmt.Sprintf("%s", userLogin)

	ctx, cancel := context.WithTimeout(c.Request().Context(), ctxTime)
	defer cancel()

	res, err := b.birdService.Create(ctx, bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "create",
		}).Errorf("unable to create bird %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, res)
}

// Get func for echo request.
func (b *BirdHandler) Get(c echo.Context) error {

	userLogin := c.Get("Login")

	ctx, cancel := context.WithTimeout(c.Request().Context(), ctxTime)
	defer cancel()

	res, err := b.birdService.Get(ctx, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "get",
		}).Errorf("unable to get bird %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// Change func for echo request.
func (b *BirdHandler) Change(c echo.Context) error {

	userLogin := c.Get("Login")

	bird := models.Bird{}

	err := c.Bind(&bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "bind model",
		}).Errorf("dogHandler: unable to bind dog data %v,", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(bird)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bird.Owner = fmt.Sprintf("%s", userLogin)

	ctx, cancel := context.WithTimeout(c.Request().Context(), ctxTime)
	defer cancel()

	res, err := b.birdService.Change(ctx, bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "change",
		}).Errorf("unable to change bird %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, res)
}

// Delete func for echo request.
func (b *BirdHandler) Delete(c echo.Context) error {

	userLogin := c.Get("Login")

	ctx, cancel := context.WithTimeout(c.Request().Context(), ctxTime)
	defer cancel()

	res, err := b.birdService.Delete(ctx, fmt.Sprintf("%s", userLogin))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "delete",
		}).Errorf("unable to delete bird %v", err)

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, res)
}

// NewBirdHandler create new handler for echo.
func NewBirdHandler(birdService *service.BirdService) *BirdHandler {
	return &BirdHandler{birdService: birdService}
}
