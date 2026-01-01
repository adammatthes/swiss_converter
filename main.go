package main

import (
	"net/http"
	"github.com/adammatthes/swiss_converter/internal/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/home", handlers.ConversionMenu)
	http.HandleFunc("/api/get-conversion-options", handlers.GenerateTargetOptions)
	http.ListenAndServe(":8080", nil)
}
