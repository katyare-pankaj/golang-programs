package main

import (
	"fmt"
	"net/http"
	"os"
)

// HTTPStatus contains information about each HTTP status code
type HTTPStatus struct {
	Code        int
	Description string
	Message     string
}

// GenerateHTTPStatusDocumentation creates a formatted string of the HTTP status documentation
func GenerateHTTPStatusDocumentation(statuses []HTTPStatus) string {
	var documentation string
	documentation += "HTTP Status Codes:\n"
	documentation += "===================\n"

	for _, status := range statuses {
		documentation += fmt.Sprintf("%d - %s\n", status.Code, status.Description)
		documentation += fmt.Sprintf("    - %s\n", status.Message)
		documentation += "\n"
	}

	return documentation
}

// GetHTTPStatuses retrieves a list of HTTP statuses
func GetHTTPStatuses() []HTTPStatus {
	var statuses []HTTPStatus

	for code, message := range http.StatusText {
		switch code {
		case http.StatusContinue, http.StatusSwitchingProtocols:
			statuses = append(statuses, HTTPStatus{code, "Informational", message})
		case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInformation,
			http.StatusNoContent, http.StatusResetContent, http.StatusPartialContent:
			statuses = append(statuses, HTTPStatus{code, "Successful", message})
		case http.StatusMultipleChoices, http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther,
			http.StatusNotModified, http.StatusUseProxy, http.StatusSwitchProxy, http.StatusTemporaryRedirect:
			statuses = append(statuses, HTTPStatus{code, "Redirection", message})
		case http.StatusBadRequest, http.StatusUnauthorized, http.StatusPaymentRequired, http.StatusForbidden,
			http.StatusNotFound, http.StatusMethodNotAllowed, http.StatusNotAcceptable, http.StatusProxyAuthenticationRequired,
			http.StatusRequestTimeout, http.StatusConflict, http.StatusGone, http.StatusLengthRequired,
			http.StatusPreconditionFailed, http.StatusPayloadTooLarge, http.StatusURITooLong, http.StatusUnsupportedMediaType,
			http.StatusRequestedRangeNotSatisfiable, http.StatusExpectationFailed, http.StatusImATeapot,
			http.StatusMisdirectedRequest, http.StatusUnprocessableEntity, http.StatusLocked, http.StatusFailedDependency,
			http.StatusUpgradeRequired, http.StatusPreconditionRequired, http.StatusTooManyRequests, http.StatusHTTPVersionNotSupported,
			http.StatusVariantAlsoNegotiates, http.StatusInsufficientStorage, http.StatusLoopDetected,
			http.StatusNotExtended, http.StatusNetworkAuthenticationRequired:
			statuses = append(statuses, HTTPStatus{code, "Client Error", message})
		case http.StatusInternalServerError, http.StatusNotImplemented, http.StatusBadGateway, http.StatusServiceUnavailable,
			http.StatusGatewayTimeout, http.StatusHTTPVersionNotSupported:
			statuses = append(statuses, HTTPStatus{code, "Server Error", message})
		}
	}

	return statuses
}

func main() {
	// Generate documentation
	statuses := GetHTTPStatuses()
	documentation := GenerateHTTPStatusDocumentation(statuses)

	// Save documentation to a file
	file, err := os.Create("http_status_codes.md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(documentation)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTTP Status Code documentation generated and saved to 'http_status_codes.md'.")
}
