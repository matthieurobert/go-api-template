package config

var ENV = &Env{}

var POSTGRES = &PostgresServer{}

func Init() {
	ENV.Initenv()

	POSTGRES.ConnectToDB(*ENV)
}
