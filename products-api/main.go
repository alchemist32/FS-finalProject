package main

import (
	"fmt"

	"github.com/products-api/core/infrastructure"
)

func main() {
	fmt.Println("Hello world!!")
	r := infrastructure.Setup()
	fmt.Println("products api: listen on port 9000")
	err := r.Run("localhost:9000")
	if r != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}
