package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"time"
)

// DogHandler creates new dog handler.
type DogHandler struct {
	dogService *service.DogService
}

// Create func for echo request.
func (h *DogHandler) Create(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Println(err)
	}

	err = c.Validate(dog)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

	err = h.dogService.Create(ctx, dog)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}
	return &echo.HTTPError{
		Code:     http.StatusOK,
		Message:  "Dog created!",
		Internal: nil,
	}
}

// Get func for echo request.
func (h *DogHandler) Get(c echo.Context) error {
	name := c.QueryParam("name")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	resultDog, err := h.dogService.Get(ctx, name)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}
	return &echo.HTTPError{
		Code:     http.StatusOK,
		Message:  resultDog,
		Internal: nil,
	}
}

// Change func for echo request.
func (h *DogHandler) Change(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Println(err)
	}

	name := c.QueryParam("name")

	// validate dog data
	err = c.Validate(dog)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	// check is dog exist
	_, err = h.dogService.Get(ctx, name)
	if err != nil {
		log.Println(err)
	}

	err = h.dogService.Change(ctx, name, dog)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}

	return &echo.HTTPError{
		Code:     http.StatusOK,
		Message:  "Dog was changed!",
		Internal: nil,
	}
}

// Delete func for echo request.
func (h *DogHandler) Delete(c echo.Context) error {
	name := c.QueryParam("name")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	// check is dog exist
	_, err := h.dogService.Get(ctx, name)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "Dog doesnt exist!",
			Internal: nil,
		}
	}

	err = h.dogService.Delete(ctx, name)
	if err != nil {
		log.Println(err)
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: nil,
		}
	}

	return &echo.HTTPError{
		Code:     http.StatusOK,
		Message:  "Dog was deleted!",
		Internal: nil,
	}
}

// NewDogHanlder create new handler for echo.
func NewDogHanlder(dogService *service.DogService) DogHandler {
	return DogHandler{dogService: dogService}
}
