package main

import (
	"fmt"

	"github.com/matthieurobert/go-api-template/config"
)

func main() {
	config.Init()

	fmt.Println("It works !!")
}
