package statusHTTP

import (
	"encoding/json"
	"testing"
	"errors"
	"net/http/httptest"	
)

func DefaultTest(t *testing.T) {
	rqrs := httptest.NewRecorder()

	var statusHTTPResponse statusHTTPError
	statusCode := 500
	errorType := statusCodes[statusCode]
	message := errorType

	statusHTTP(rqrs, statusCode)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Not a valid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("wrong status code!: got %v need %v", rqrs.Code, statusCode)
	}
}

func TestMessage(t *testing.T) {
	rqrs := httptest.NewRecorder()

	var statusHTTPResponse statusHTTPError
	statusCode := 500
	errorType := statusCodes[statusCode]
	message := "Some random msg for test"

	statusHTTP(rqrs, statusCode, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v need %v", rqrs.Code, statusCode)
	}
}

func TestErrorMessage(t *testing.T) {
	rqrs := httptest.NewRecorder()

	var statusHTTPResponse statusHTTPError
	statusCode := 500
	errorType := statusCodes[statusCode]
	message := errors.New("Random msg for error")

	statusHTTP(rqrs, statusCode, message)

	if err := json.Unmarshal(rqrs.Body.Bytes(), &statusHTTPResponse); err != nil {
		t.Errorf("Invalid Response!: %v", err)
	}

	if statusHTTPResponse.ErrorType != errorType || statusHTTPResponse.Message != message.Error() || statusHTTPResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rqrs.Code != statusCode {
		t.Errorf("Wrong status code: got %v need %v", rqrs.Code, statusCode)
	}
}