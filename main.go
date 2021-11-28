package main

import (
	"fmt"
	"net/http"

	"github.com/matthieurobert/go-api-template/config"
)

func main() {
	config.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

	http.ListenAndServe(":8000", nil)

	fmt.Println("It works !!")
}
