package initializers

import (
	"errors"
	// "fmt"
	"os"
)

type Config struct {
	username    string
	password    string
	port        string
	host        string
	name        string
	appname     string
	Port_server string

	ClientOrigin string

	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	AccessTokenExpiresIn   string
	RefreshTokenExpiresIn  string
	AccessTokenMaxAge      string
	RefreshTokenMaxAge     string
}

func LoadEnv() (Config, error) {
	config := Config{
		username:    os.Getenv("DB_USERNAME"),
		password:    os.Getenv("DB_PASSWORD"),
		port:        os.Getenv("DB_PORT"),
		host:        os.Getenv("DB_HOST"),
		name:        os.Getenv("DB_NAME"),
		appname:     os.Getenv("APP_NAME"),
		Port_server: os.Getenv("SERVER_PORT"),


		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),

		AccessTokenPrivateKey:  os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"),
		AccessTokenPublicKey:   os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"),
		RefreshTokenPrivateKey: os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"),
		RefreshTokenPublicKey:  os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"),
		AccessTokenExpiresIn:   os.Getenv("ACCESS_TOKEN_EXPIRED_IN"),
		RefreshTokenExpiresIn:  os.Getenv("REFRESH_TOKEN_EXPIRED_IN"),
		AccessTokenMaxAge:      os.Getenv("ACCESS_TOKEN_MAXAGE"),
		RefreshTokenMaxAge:     os.Getenv("REFRESH_TOKEN_MAXAGE"),
	}

	if config.name == "" {
		return config, errors.New("host not found")
	}
	// fmt.Println( "private key : "+config.AccessTokenPrivateKey)

	return config, nil
}
