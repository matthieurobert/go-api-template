package config

import "github.com/sirupsen/logrus"

// ENV is a global variable knowing needed environnement variable
var ENV = &Env{}

// POSTGRES is a global variable for postgres server
var POSTGRES = &PostgresServer{}

// LOGGER is a global variable for logger object from logrus librairy
var LOGGER = &logrus.Logger{}

// InitConfig initiates the config
func InitConfig() {
	// Get environnement variables
	ENV.Initenv()

	// Init Logger
	LOGGER = initLogger(*ENV)

	// Open connection to postgres database
	POSTGRES.ConnectToDB(*ENV)
}
