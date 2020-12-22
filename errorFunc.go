package statusHTTP

import (
	"net/http"
)

// Unauthorized for 401
func Unauthorized(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 401, message...)
}

// Forbidden for 403
func Forbidden(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 403, message...)
}

// BadRequest for 400
func BadRequest(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 400, message...)
}

// Internal Server Error for 500
func Internal(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 500, message...)
}

// NotFound for 404
func NotFound(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 404, message...)
}

// BadGateway for 502
func BadGateway(w http.ResponseWriter, message ...interface{}) {
	statusHTTP(w, 502, message...)
}

