package main

import (
	"database/sql"
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
)

type server struct {
	dogRepo repository.DogRepository
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
	}

	base, err := sql.Open("postgres", "port="+cfg.Port+" host="+cfg.Host+" user="+cfg.User+" password="+cfg.Password+" dbname="+cfg.Dbname+" sslmode="+cfg.Sslmode)
	if err != nil {
		log.Fatal(err)
	}
	if err := base.Ping(); err != nil {
		log.Println(err)
	}
	s := server{}
	s.dogRepo = repository.NewDogRepository(base)
	dogService := service.NewDogService(s.dogRepo)

	dogHabdler := handler.NewDogHanlder(dogService)

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}

	e.PUT("/dogs/", dogHabdler.Create)
	e.GET("/dogs/", dogHabdler.Get)
	e.POST("/dogs/", dogHabdler.Change)
	e.DELETE("/dogs/", dogHabdler.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
