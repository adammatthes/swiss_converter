package main

import (
	"net/http"
	"embed"
	"log"
	"io/fs"
	"fmt"
	"github.com/adammatthes/swiss_converter/internal/handlers"
	"github.com/adammatthes/swiss_converter/internal/database"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	endpoint := http.FileServer(http.FS(staticFS))


	err = fs.WalkDir(staticFiles, ".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fmt.Printf("Embedded file found: %s\n", path)
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	db := database.SetupDatabase()

	mux := http.NewServeMux()


	mux.Handle("/static/", http.StripPrefix("/static/", endpoint))

	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/", handlers.ServeIndexPage)
	mux.HandleFunc("/favicon.ico", handlers.ServeFavicon)
	mux.HandleFunc("/home", handlers.ConversionMenu)

	mux.HandleFunc("/api/get-conversion-options", handlers.GenerateTargetOptions)
	mux.HandleFunc("/api/get-starting-types", handlers.GenerateStartingOptions)
	mux.HandleFunc("/api/convert", handlers.ProcessConversion)

	log.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

	db.Close()
}
