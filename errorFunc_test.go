 package statusHTTP

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestForUnathorized(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 401
	errorType := "Unauthorized"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	Unauthorized(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

func TestForForbidden(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 403
	errorType := "Forbidden"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	Forbidden(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

func TestForBadRequest(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 400
	errorType := "Bad Request"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	BadRequest(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

func TestForInternalServerError(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 500
	errorType := "Internal Server Error"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	Internal(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

func TestForNotFound(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 404
	errorType := "Not Found"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	NotFound(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

func TestForBadGateway(t *testing.T) {
	rqrs := httptest.NewRecorder()

	statusCode := 502
	errorType := "Bad Gateway"
	message := "Random msg for Test"

	var statusHTTPResponse statusHTTPError

	BadGateway(rqrs, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v want %v", rqrs.Code, statusCode)
	}
}

