package server

import (
	"Dogs/src/internal/handler"
	"Dogs/src/internal/repository"
	"Dogs/src/internal/service"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

type Server struct {
	dogRepo repository.DogPostgresRepository
}

func Start()  {


	base, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := base.Ping(); err != nil {
		log.Println(err)
	}
	s := Server{}
	s.dogRepo = repository.NewDogPostgresRepository(base)
	dogService := service.NewDogService(s.dogRepo)

	dogHabdler := handler.NewDogHanlder(dogService)

	e := echo.New()

	e.PUT("/dogs/",dogHabdler.CreateDog)


	e.Logger.Fatal(e.Start(":1323"))
}
