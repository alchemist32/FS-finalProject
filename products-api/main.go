package main

import (
	"fmt"

	"github.com/products-api/core/infrastructure"
	"github.com/products-api/core/infrastructure/config"
)

func main() {
	config.LoadEnv()
	port := config.ServerPort
	r := infrastructure.Setup()
	fmt.Println("products api: listen on port" + port)
	err := r.Run(":" + port)
	if r != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}
