package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type serverConfig struct { // prefix SERVER_
	Address string
}
type databaseConfig struct { // prefix DATABASE_
	Username     string
	Password     string
	Host         string
	DatabaseName string
}
type config struct {
	Server   serverConfig
	Database databaseConfig
}

var Config = config{}

func (this *config) Init() {
	godotenv.Load()
	godotenv.Load("dev.env")

	this.Server = serverConfig{
		Address: os.Getenv("SERVER_ADDRESS"),
	}

	this.Database = databaseConfig{
		Username:     os.Getenv("DATABASE_USERNAME"),
		Password:     os.Getenv("DATABASE_PASSWORD"),
		Host:         os.Getenv("DATABASE_HOST"),
		DatabaseName: os.Getenv("DATABASE_DATABASENAME"),
	}
}
