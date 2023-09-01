package initializers

import (
	"errors"
	"os"
)

type Config struct {
	username    string
	password    string
	port        string
	host        string
	name        string
	appname     string
	port_server string
}

func LoadEnv() (Config,error) {
	config := Config{
		username:    os.Getenv("DB_USERNAME"),
		password:    os.Getenv("DB_PASSWORD"),
		port:        os.Getenv("DB_PORT"),
		host:        os.Getenv("DB_HOST"),
		name:        os.Getenv("DB_NAME"),
		appname:     os.Getenv("APP_NAME"),
		port_server: os.Getenv("SERVER_PORT"),
	}

	if config.name == ""  {
		return config, errors.New("host not found")
	} 


	return config, nil
}
