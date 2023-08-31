package config

import (
	"final-project-prakerja/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDatabase()
}

type Config struct {
	username string
	password string
	port     string
	host     string
	name     string
}

func InitDatabase() {
	config := Config{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		port:     os.Getenv("DB_PORT"),
		host:     os.Getenv("DB_HOST"),
		name:     os.Getenv("DB_NAME"),
	}

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.username,
		config.password,
		config.host,
		config.port,
		config.name,
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	DB = db
	if err != nil {
		panic(err)
	}
	InitialMmigrate()
}

func InitialMmigrate()  {
	DB.AutoMigrate(&models.User{})
}