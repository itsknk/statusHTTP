package statusHTTP

import (
	"net/http"
)

//Unauthorized
func Unauthorized(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 401, message...)
}

//Forbidden
func Forbidden(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 403, message...)
}

//BadRequest
func BadRequest(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 400, message...)
}

//Internal Server Error
func Internal(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 500, message...)
}

//NotFound
func NotFound(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 404, message...)
}

//BadGateway
func BadGateway(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 502, message...)
}

