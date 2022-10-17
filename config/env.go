package config

import (
	"os"
	"strconv"
)

// Env : structure knowing all environnement variable
type Env struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	ApiPort          int
	LogLevel         string
}

// InitEnv : method setting values to environnement variables
func (env *Env) Initenv() {
	var err error

	env.PostgresHost = os.Getenv("POSTGRES_HOST")
	env.PostgresUser = os.Getenv("POSTGRES_USER")
	env.PostgresPort, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if err != nil {
		panic(err)
	}

	env.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	env.PostgresDatabase = os.Getenv("POSTGRES_DB")
	env.ApiPort, _ = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		panic(err)
	}

	env.LogLevel = os.Getenv("LOG_LEVEL")
}
