package main

import (
	"final-project-prakerja/initializers"
	"final-project-prakerja/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadEnv()

	if err != nil {
		log.Fatal("🚀 Could not load environment variables")
	}

	initializers.ConnectDB(&config)

}

func main() {

	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("error mogration ", err)
	} else {
		fmt.Println("👍 Migration complete")
	}
}
