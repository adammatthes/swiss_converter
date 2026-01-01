package main

import (
	"net/http"
	"embed"
	"github.com/adammatthes/swiss_converter/internal/handlers"
)

var staticFiles embed.FS

func main() {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/favicon.ico", handlers.ServeFavicon)
	http.HandleFunc("/home", handlers.ConversionMenu)
	http.HandleFunc("/api/get-conversion-options", handlers.GenerateTargetOptions)
	http.ListenAndServe(":8080", nil)
}
