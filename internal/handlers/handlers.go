package handlers

import (
	"io"
	"fmt"
	"encoding/json"
	"strings"
	"net/http"
)

type UserRequest struct {
	Value string `json:"name"`
	Type string `json:"type"`
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello, world!</h1>")
}

func generateDropdownOptions(options []string) string {
	var result []string

	for _, option := range options {
		nextOpt := fmt.Sprintf("<option class=\"dropdownOptions\">%s</option>", option)
		result = append(result, nextOpt)
	}

	return strings.Join(result, "\n")
}

func ConversionMenu(w http.ResponseWriter, req *http.Request) {
	startingOptions := []string{"Hexadecimal", "Decimal", "Binary"}

	htmlOptions := generateDropdownOptions(startingOptions)

	firstDropdown := fmt.Sprintf("<select class=\"dropdownMenu\">%s</select>", htmlOptions)

	io.WriteString(w, fmt.Sprintf(`<div class="conversionMenu">%s</div>`, firstDropdown))
}

func GenerateTargetOptions(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := map[string][]string{"options": []string{"Hexadecimal", "Decimal", "Binary"}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
