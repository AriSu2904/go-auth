package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, status int, code, message string) {
	errorResponse := map[string]map[string]string{
		"error": {
			"code":    code,
			"message": message,
		},
	}
	WriteJSON(w, status, errorResponse)
}

func WriteValidationError(w http.ResponseWriter, err error) {
	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessage := "Field '" + err.Field() + "' failed on the '" + err.Tag() + "' tag"
		errorMessages = append(errorMessages, errorMessage)
	}
	WriteError(w, http.StatusBadRequest, "BAD_REQUEST",
		"Invalid request body: "+strings.Join(errorMessages, "; "))
}
