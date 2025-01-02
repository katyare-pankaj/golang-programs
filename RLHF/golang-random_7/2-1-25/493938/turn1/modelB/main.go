package main

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrInvalidRequest   ErrorCode = "INVALID_REQUEST"
	ErrInternalServer   ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrResourceNotFound ErrorCode = "RESOURCE_NOT_FOUND"
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewAppError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = NewAppError(ErrInternalServer, err.Error())
	}

	// Customize error message based on user context or system state
	userContext := r.Context().Value("user_context").(string)
	if userContext == "admin" {
		appErr.Message = fmt.Sprintf("%s: %s (user context: %s)", appErr.Code, appErr.Message, userContext)
	}

	// Customize error message for multilingual support
	lang := r.Header.Get("Accept-Language")
	switch lang {
	case "fr":
		switch appErr.Code {
		case ErrInvalidRequest:
			appErr.Message = "Requête invalide"
		case ErrInternalServer:
			appErr.Message = "Erreur interne du serveur"
		case ErrResourceNotFound:
			appErr.Message = "Ressource non trouvée"
		}
	default:
		// Use the default message in English
	}

	http.Error(w, appErr.Message, http.StatusInternalServerError)
}

func main() {
	// Sample usage:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulate an error
		err := NewAppError(ErrInvalidRequest, "Invalid query parameter")
		handleError(w, r, err)
	})

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
