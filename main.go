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
	queries := database.New(db)

	app := handlers.Application{Db : db, Queries: queries}

	mux := http.NewServeMux()


	mux.Handle("/static/", http.StripPrefix("/static/", endpoint))

	mux.HandleFunc("/hello", app.HelloHandler)
	mux.HandleFunc("/", app.ServeIndexPage)
	mux.HandleFunc("/favicon.ico", app.ServeFavicon)
	mux.HandleFunc("/home", app.ConversionMenu)

	mux.HandleFunc("/api/get-conversion-options", app.GenerateTargetOptions)
	mux.HandleFunc("/api/get-starting-types", app.GenerateStartingOptions)
	mux.HandleFunc("/api/convert", app.ProcessConversion)
	mux.HandleFunc("/api/currency", app.ProcessCurrency)
	mux.HandleFunc("/api/create-conversion", app.CreateConversion)
	mux.HandleFunc("/api/delete-conversion", app.DeleteConversion)
	mux.HandleFunc("/api/metrics", app.GetMetrics)

	log.Println("\nServer starting on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

	db.Close()
}
