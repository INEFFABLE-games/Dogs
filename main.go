package main

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
)

type server struct {
	dogRepo  repository.DogRepository
	userRepo repository.UserRepository
}

func newServer(dogRepo *repository.DogRepository, userRepo *repository.UserRepository) *server {
	return &server{
		dogRepo:  *dogRepo,
		userRepo: *userRepo,
	}
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
	log.SetLevel(log.DebugLevel)
	cfg := config.NewConfig()

	base, err := sql.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Port,
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Sslmode))
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	if err := base.Ping(); err != nil {
		log.Errorf("main: unable to ping sql connection %v,", err)
	}

	s := newServer(repository.NewDogRepository(base), repository.NewUserRepository(base))
	dogService := service.NewDogService(s.dogRepo)
	dogHandler := handler.NewDogHandler(dogService)

	userService := service.NewUserService(s.userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}

	dogsPath := "/dogs/"
	usersPath := "/users/"

	//dog routs
	e.PUT(dogsPath, dogHandler.Create)
	e.GET(dogsPath, dogHandler.Get)
	e.POST(dogsPath, dogHandler.Change)
	e.DELETE(dogsPath, dogHandler.Delete)

	// user routs
	e.PUT("/users/registration/", userHandler.Create)
	e.PUT(usersPath, userHandler.Login)
	e.DELETE(usersPath, userHandler.Logout)

	e.Logger.Fatal(e.Start(":1323"))
}
