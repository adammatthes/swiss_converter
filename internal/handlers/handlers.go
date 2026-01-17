package handlers

import (
	"io"
	"fmt"
	"encoding/json"
	"strings"
	"net/http"
	"database/sql"
	"strconv"
	"context"
	"sort"
	"slices"
	"github.com/adammatthes/swiss_converter/internal/database"
	"github.com/adammatthes/swiss_converter/internal/conversion_options"
	"github.com/adammatthes/swiss_converter/internal/convert"
)

type Application struct {
	Db	*sql.DB
	Queries	*database.Queries
}

type UserRequest struct {
	Value string `json:"name"`
	Type string `json:"type"`
	Category string `json:"category"`
}

type ConversionRequest struct {
	Category string `json:"category"`
	StartType string `json:"start-type"`
	EndType string `json:"end-type"`
	Value	string `json:"value"`
}

type NewConversion struct {
	StartType 	string `json:"start-type"`
	EndType		string `json:"end-type"`
	Value 		string `json:"value"`
}

func (app *Application) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello, world!</h1>")
}

func (app *Application) ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func (app *Application) ServeFavicon(w http.ResponseWriter, r *http.Request) {
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

func generateButton(buttonId, innerFieldsId, buttonText string) string {
	return fmt.Sprintf(`<button id="%s">%s</button><div id="%s"></div>`, buttonId, buttonText, innerFieldsId)
}

func (app *Application) ConversionMenu(w http.ResponseWriter, req *http.Request) {
	startingOptions := []string{conversion_options.Base,
				conversion_options.Distance,
				conversion_options.Currency,
				conversion_options.Temperature,
				conversion_options.Custom,
			}

	htmlOptions := generateDropdownOptions(startingOptions)

	firstDropdown := fmt.Sprintf(`<select id="categorySelect" class="dropdownMenu">%s</select>`, htmlOptions)

	conversionGenerateButton := generateButton("customRateInitiator", "customRateFields", "Add a New Conversion")
	metricsButton := generateButton("metricsCreateButton", "metricsTable", "Conversion Metrics")

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
	<div id="resultOutputSection"></div>
	<div id="customGenerationSection">%s</div>
	<div id="metricsSection">%s</div>
	<script src="./static/script.js"></script>
	</body>
	</html>
	`, firstDropdown, conversionGenerateButton, metricsButton))
}

func (app *Application) GenerateStartingOptions(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	var options []string

	if req.Type == "Custom" {
		fmt.Println("Custom category detected")
		options, err = app.Queries.GetStartingCustomOptions(r.Context())
		fmt.Printf("%v", options)
		if err != nil {
			fmt.Println("Didn't find Custom Starting Options: %v", err)
		}
	} else {
		options, err = conversion_options.GetTypesByCategory(req.Type)
	}

	response := map[string][]string{"options": options}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *Application) GenerateTargetOptions(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var options []string
	if (req.Category == "Custom") {
		options, err = app.Queries.GetCustomConversionOptions(r.Context(), req.Type)
		if err != nil {
			http.Error(w, fmt.Sprint("Error getting Custom target conversion options: %v", err), http.StatusBadRequest)
			return
		}
	} else {
		options, err = conversion_options.GetConversionOptions(req.Type)
	}

	response := map[string][]string{"options": options}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *Application) ProcessConversion(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if (req.StartType == "Select an option" || req.EndType == "Select an option") {
		http.Error(w, "Missing starting or ending conversion type", http.StatusBadRequest)
		return
	}

	var result string

	if req.Category == "Custom" {
		startVal, err := strconv.ParseFloat(req.Value, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		params := database.GetCustomExchangeRateParams{StartType: req.StartType, EndType: req.EndType}
		exchangeRate, err := app.Queries.GetCustomExchangeRate(r.Context(), params)
		result = fmt.Sprintf("%v", startVal * exchangeRate)
	} else {
		function, err := convert.GetConversionFunction(req.StartType, req.EndType)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		result, err = function(req.Value)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}
	}


	response := map[string]string{"result": fmt.Sprintf("%v", result)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	go func(st, et, val string) {
		queryLogParams := database.AddQueryLogEntryParams{StartType: st, EndType: et, Amount: val}
		_ = app.Queries.AddQueryLogEntry(context.Background(), queryLogParams)
	}(req.StartType, req.EndType, req.Value)
}

func (app *Application) ProcessCurrency(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	start := req.StartType
	end := req.EndType

	exchangeRate, err := app.Queries.GetExchangeRate(r.Context(), start+end)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	startVal, err := strconv.ParseFloat(req.Value, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	response := map[string]string{"result": fmt.Sprintf("%v", startVal * exchangeRate)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	go func(st, et, val string) {
		queryLogParams := database.AddQueryLogEntryParams{StartType: st, EndType: et, Amount: val}
		_ = app.Queries.AddQueryLogEntry(context.Background(), queryLogParams)
	}(req.StartType, req.EndType, req.Value)
}

func (app *Application) GetCustomOptions(w http.ResponseWriter, r *http.Request) ([]string) {
	options, err := app.Queries.GetStartingCustomOptions(r.Context())
	if err != nil {
		fmt.Errorf("Error getting Starting Custom Options: %v", err)
		return []string{"No Custom Options Found", fmt.Sprintf("%v", err)}
	}

	return options
}

func (app *Application) CreateConversion(w http.ResponseWriter, r *http.Request) {
	var req NewConversion
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	start := req.StartType
	end := req.EndType
	rate, err := strconv.ParseFloat(req.Value, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	gcerp := database.GetCustomExchangeRateParams{StartType: start, EndType: end}

	checkIfExists, err := app.Queries.GetCustomExchangeRate(r.Context(), gcerp)
	if err == nil || checkIfExists != 0.0 {
		ucep := database.UpdateCustomExchangeParams{ExchangeRate: rate, StartType: start, EndType: end}
		err = app.Queries.UpdateCustomExchange(r.Context(), ucep)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
	} else {
		accp := database.AddCustomConversionParams{StartType: start, EndType: end, ExchangeRate: rate}
		err = app.Queries.AddCustomConversion(r.Context(), accp)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}
	}

	err = app.DeduceNewConversions(start, end, rate)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]any{"success":true, "error": err}
	json.NewEncoder(w).Encode(response)
}

func (app *Application) DeduceNewConversions(start_type, end_type string, rate float64) error {
	inverse_rate := 1.0 / rate
	gcerParams := database.GetCustomExchangeRateParams{StartType: end_type, EndType: start_type}

	_, err := app.Queries.GetCustomExchangeRate(context.Background(), gcerParams)
	if err != nil {
		if err == sql.ErrNoRows {
			accParams := database.AddCustomConversionParams{StartType: end_type, EndType: start_type, ExchangeRate: inverse_rate}
			err = app.Queries.AddCustomConversion(context.Background(), accParams)
			if err != nil {
				return fmt.Errorf("Error adding inverse rate: %v", err)

			}
		} else {
			return fmt.Errorf("Database Error: %v", err)
		}
	} else {
		uceParams := database.UpdateCustomExchangeParams{ExchangeRate: inverse_rate, StartType: end_type, EndType: start_type}
		err = app.Queries.UpdateCustomExchange(context.Background(), uceParams)
		if err != nil {
			return fmt.Errorf("Error updating inverse rate: %v", err)
		}
	}

	onceRemoved, err := app.Queries.GetCustomConversionOptions(context.Background(), end_type)
	if err != nil {
		return fmt.Errorf("Error getting once-removed conversions: %v", err)
	}

	for _, opt := range onceRemoved {
		tempParams := database.GetCustomExchangeRateParams{StartType: end_type, EndType: opt}
		secondRate, err := app.Queries.GetCustomExchangeRate(context.Background(), tempParams)

		checkParams := database.GetCustomExchangeRateParams{StartType: start_type, EndType: opt}
		_, err = app.Queries.GetCustomExchangeRate(context.Background(), checkParams)
		if err != nil {
			if err == sql.ErrNoRows {
				accParams := database.AddCustomConversionParams{StartType: start_type, EndType: opt, ExchangeRate: rate * secondRate}
				err = app.Queries.AddCustomConversion(context.Background(), accParams)
				if err != nil {
					return fmt.Errorf("Error adding once removed conversion: %v", err)
				}
			} else {
				return fmt.Errorf("Database Error: %v", err)
			}
		} else {
			uceParams := database.UpdateCustomExchangeParams{ExchangeRate: rate * secondRate, StartType: start_type, EndType: opt}
			err = app.Queries.UpdateCustomExchange(context.Background(), uceParams)
			if err != nil {
				return fmt.Errorf("Error updating once removed conversion: %v", err)
			}
		}

	}

	return nil
}

func (app *Application) GetMetrics(w http.ResponseWriter, r *http.Request) {
	type countObject struct {
		StartType	string `json:start_type`
		EndType		string `json:end_type`
		Count		uint64 `json:count`
	}

	bakedTypes := conversion_options.GetAllTypes()

	customTypes, _ := app.Queries.GetStartingCustomOptions(r.Context())

	combined := append(bakedTypes, customTypes...)
	sort.Strings(combined)

	allTypes := slices.Compact(combined)

	var result []countObject

	for _, t1 := range allTypes {
		for _, t2 := range allTypes {
			if t1 == t2 {
				continue
			}

			gccParams := database.GetConversionCountParams{StartType: t1, EndType: t2}
			count, err := app.Queries.GetConversionCount(r.Context(), gccParams)
			if err != nil || count == 0 {
				continue
			}

			result = append(result, countObject{StartType: t1, EndType: t2, Count: uint64(count)})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
