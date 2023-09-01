package main

import (
	"final-project-prakerja/initializers"
	"fmt"
)

func main() {
	config := initializers.LoadEnv()
	fmt.Println(config)
	initializers.ConnectDB(&config)
}