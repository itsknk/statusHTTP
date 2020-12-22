package statusHTTP

import (
	"encoding/json"
	"net/http"
)

var statusCodes = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
	502: "Bad Gateway",
}

type statusHTTPError struct {
	ErrorType  string `json:"error"`
	Message    string `json:"message`
	StatusCode int    `json:"statusCode"`
}

func statusHTTP(w http.ResponseWriter, statusCode int, args ...interface{}) {
	errorType := statusCodes[statusCode]
	message := errorType
	if len(args) != 0 {
		switch args[0].(type) {
		case error:
			message = args[0].(error).Error()
		case string:
			message = args[0].(string)
		}
	}
	errorString, _ := json.Marshal(statusHTTPError{
		errorType,
		message,
		statusCode,
	})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(errorString)
}
