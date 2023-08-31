package main

import (
	"final-project-prakerja/config"
	"fmt"
)

func main() {
	fmt.Println("test")
	db :=  config.DB.Migrator().CurrentDatabase()
	fmt.Println(db)
}