package handlers

import (
	"io"
	"fmt"
	"encoding/json"
	"strings"
	"net/http"
	"github.com/adammatthes/swiss_converter/internal/conversion_options"
)

type UserRequest struct {
	Value string `json:"name"`
	Type string `json:"type"`
}

type ConversionRequest struct {
	Category string `json:"category"`
	StartType string `json:"start-type"`
	EndType string `json:"end-type"`
	Value	string `json:"value"`
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello, world!</h1>")
}

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ServeFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	http.ServeFile(w, r, "./static/favicon.ico")
}

func generateDropdownOptions(options []string) string {
	var result []string

	result = append(result, `<option disabled selected class="dropdownOptions">Select an option</option>`)
	for _, option := range options {
		nextOpt := fmt.Sprintf("<option class=\"dropdownOptions\">%s</option>", option)
		result = append(result, nextOpt)
	}

	return strings.Join(result, "\n")
}

func ConversionMenu(w http.ResponseWriter, req *http.Request) {
	startingOptions := []string{conversion_options.Base, conversion_options.Distance}

	htmlOptions := generateDropdownOptions(startingOptions)

	firstDropdown := fmt.Sprintf(`<select id="categorySelect" class=\"dropdownMenu\">%s</select>`, htmlOptions)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Swiss Converter</title>
		<link rel="stylesheet" href="/static/style.css">
	</head>
	<body>
	<div id="conversionMenu">%s</div>
	<script src="./static/script.js"></script>
	</body>
	</html>
	`, firstDropdown))
}

func GenerateStartingOptions(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	options, err := conversion_options.GetTypesByCategory(req.Type)

	response := map[string][]string{"options": options}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GenerateTargetOptions(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	options, err := conversion_options.GetConversionOptions(req.Type)

	response := map[string][]string{"options": options}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ProcessConversion(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := map[string]string{"result": fmt.Sprintf("%v", req.Value)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
