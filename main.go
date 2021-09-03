package main

import (
	"Dogs/internal/Validators"
	"Dogs/internal/config"
	"Dogs/internal/handler"
	"Dogs/internal/middleware"
	"Dogs/internal/models"
	"Dogs/internal/repository"
	"Dogs/internal/service"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	log.SetLevel(log.DebugLevel)
	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Port,
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Sslmode))
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	if err := conn.Ping(); err != nil {
		log.Errorf("main: unable to ping sql connection %v,", err)
	}

	grpcConn, err := grpc.Dial("localhost:1323", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	//s := newServer(repository.NewDogRepository(base), repository.NewUserRepository(base))
	dogService := service.NewDogService(*repository.NewDogRepository(conn))
	dogHandler := handler.NewDogHandler(dogService)

	birdService := service.NewBirdService(grpcConn)
	birdHandler := handler.NewBirdHandler(birdService)

	userService := service.NewUserService(*repository.NewUserRepository(conn))
	authService := service.NewAuthService(*repository.NewTokenRepository(conn))

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService, authService)

	e := echo.New()
	e.Validator = &Validators.CustomValidator{Validator: validator.New()}

	dogsPath := "/dogs/"
	usersPath := "/users/"
	birdPath := "/birds/"

	// JWT config
	jwtConfig := middleware2.JWTConfig{
		SigningKey: []byte("dog"),
		Claims:     &models.CustomClaims{},
		ContextKey: "Data",
	}

	// RT config
	rtConfig := middleware2.JWTConfig{
		SigningKey: []byte("RefTokKey"),
		Claims:     &models.CustomClaims{},
		ContextKey: "Data",
	}

	//dog routs
	e.PUT(dogsPath, dogHandler.Create, middleware.AuthenticateToken(jwtConfig))
	e.GET(dogsPath, dogHandler.Get, middleware.AuthenticateToken(jwtConfig))
	e.POST(dogsPath, dogHandler.Change, middleware.AuthenticateToken(jwtConfig))
	e.DELETE(dogsPath, dogHandler.Delete, middleware.AuthenticateToken(jwtConfig))

	//bird routs
	e.PUT(birdPath, birdHandler.Create, middleware.AuthenticateToken(jwtConfig))
	e.GET(birdPath, birdHandler.Get, middleware.AuthenticateToken(jwtConfig))
	e.POST(birdPath, birdHandler.Change, middleware.AuthenticateToken(jwtConfig))
	e.DELETE(birdPath, birdHandler.Delete, middleware.AuthenticateToken(jwtConfig))

	// user routs
	e.PUT("/users/registration/", userHandler.Create)
	e.PUT(usersPath, userHandler.Login)
	e.DELETE(usersPath, userHandler.Logout, middleware.AuthenticateToken(jwtConfig))

	// Authentication routs
	e.PUT("/users/Authentication/", authHandler.RefreshTokens, middleware.AuthenticateToken(rtConfig))

	e.Logger.Fatal(e.Start(":1333"))
}
