package config

// Config creates new database connection config from env variables.
type Config struct {
	Port     string `env:"SQLPORT,required,notEmpty"`
	Host     string `env:"SQLHOST,required,notEmpty"`
	User     string `env:"SQLUSER,required,notEmpty"`
	Password string `env:"SQLPASSWORD,required,notEmpty"`
	Dbname   string `env:"SQLDBNAME,required,notEmpty"`
	Sslmode  string `env:"SQLSSLMODE,required,notEmpty"`
}

// POSTGRES_URI = port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable
