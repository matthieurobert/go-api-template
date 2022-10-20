package entity

import (
	"os"
	"testing"

	"github.com/matthieurobert/go-api-template/config"
)

func TestMain(m *testing.M) {
	config.InitConfig()
	CreateSchema(config.POSTGRES.DB)

	exitCode := m.Run()

	config.POSTGRES.DB.Close()

	os.Exit(exitCode)
}

func TestPostUser(t *testing.T) {
	user := User{
		Username: "admin",
		Password: "admin",
	}

	userRepositoryFactory := UserRepositoryFactory{
		Database: config.POSTGRES.DB,
	}

	userRepo := userRepositoryFactory.Build()

	_, err := userRepo.PostUser(user)

	if err != nil {
		t.Errorf(err.Error())
	}
}
