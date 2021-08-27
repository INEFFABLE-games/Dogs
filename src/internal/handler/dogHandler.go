package handler

import (
	"Dogs/src/internal/models"
	"Dogs/src/internal/service"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type DogHandler struct {
	dogService *service.DogService
}

func (h *DogHandler) CreateDog(c echo.Context) error {

	dog := models.Dog{}
	err := c.Bind(&dog);if err != nil{
		log.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

	err = h.dogService.CreateDog(ctx,dog);if err != nil{
		log.Println(err)
	}

	if err != nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.String(http.StatusOK,"Dog created!")
}

func NewDogHanlder(dogService *service.DogService) DogHandler {
	return DogHandler{dogService: dogService}
}