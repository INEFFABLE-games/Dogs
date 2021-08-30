package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

// Config creates new database connection config from env variables.
type Config struct {
	Port     string `env:"SQLPORT,required,notEmpty"`
	Host     string `env:"SQLHOST,required,notEmpty"`
	User     string `env:"SQLUSER,required,notEmpty"`
	Password string `env:"SQLPASSWORD,required,notEmpty"`
	Dbname   string `env:"SQLDBNAME,required,notEmpty"`
	Sslmode  string `env:"SQLSSLMODE,required,notEmpty"`
}

// NewConfig create new config object.
func NewConfig() *Config {

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.WithFields(log.Fields{
			"handler": "config",
			"action":  "initialize",
		}).Errorf("unable to pars environment variables %v,", err)
	}

	return &cfg
}

// POSTGRES_URI = port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable
