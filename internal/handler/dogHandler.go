package handler

import (
	"Dogs/internal/models"
	"Dogs/internal/service"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type DogHandler struct {
	dogService *service.DogService
}

func (h *DogHandler) Create(c echo.Context) error {

	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Println(err)
	}

	err = c.Validate(dog)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

	err = h.dogService.Create(ctx, dog)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Dog created!")
}

func (h *DogHandler) Get(c echo.Context) error {

	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Println(err)
	}

	err = c.Validate(dog)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	resultDog, err := h.dogService.Get(ctx, dog)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resultDog)
}

func (h *DogHandler) Change(c echo.Context) error {

	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Println(err)
	}

	name := c.QueryParam("name")
	//validate dog data
	err = c.Validate(dog)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	//check is dog exist
	_, err = h.dogService.Get(ctx, models.Dog{Name: name})
	if err != nil {
		log.Println(err)
	}

	err = h.dogService.Change(ctx, name, dog)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Dog was changed!")
}

func (h *DogHandler) Delete(c echo.Context) error {

	name := c.QueryParam("name")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	//check is dog exist
	_, err := h.dogService.Get(ctx, models.Dog{Name: name})
	if err != nil {
		return c.String(http.StatusBadRequest, "Dog doesnt exist!")
	}

	err = h.dogService.Delete(ctx, name)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Dog was deleted!")
}

func NewDogHanlder(dogService *service.DogService) DogHandler {
	return DogHandler{dogService: dogService}
}
