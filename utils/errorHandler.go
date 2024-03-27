package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// To Show Log Error

func LogError(r *http.Request, err error) {

	logrus.Errorf("Error Received: %s %s %s", err, r.Method, r.URL.Path)
}

// SuccessResponse sends a JSON-formatted success response with the specified status code and data
func SuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ErrorResponse sends a JSON-formatted error response with the specified status code and message

func ErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	errResponseToSend := ErrorResponseType{
		Success: false,
		Message: message,
	}

	jsonData, err := json.Marshal(errResponseToSend)

	if err != nil {
		// If unable to marshal, log the error and send a generic error response
		LogError(r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	// Write JSON response to the ResponseWriter

	_, err = w.Write(jsonData)

	if err != nil {
		// If unable to write response, log the error
		LogError(r, err)
	}

}

// Server Error response
func ServerError(w http.ResponseWriter, r *http.Request, err error) {

	LogError(r, err)

	// Prepare a message with the error

	message := "Internal Server Error "

	// Send the error message
	ErrorResponse(w, r, http.StatusInternalServerError, message)

}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	// Prepare a message with the error
	message := fmt.Sprintf("Opps, the page you are looking for does not exist. Please check the URL: %s", path)

	LogError(r, fmt.Errorf("path not found"))

	// Send the error message
	ErrorResponse(w, r, http.StatusNotFound, message)

}

// Method Not Allowed Response

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	// Prepare a message with the error
	message := fmt.Sprintf("Method %s Not Allowed. Please check the URL: %s", r.Method, r.URL.Path)

	LogError(r, fmt.Errorf("method not allowed"))

	// Send the error message
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)

}
