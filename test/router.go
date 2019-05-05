package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Idea: https://stackoverflow.com/questions/15897803/how-can-i-have-a-common-test-suite-for-multiple-packages-in-go

// ResponseSample is one sample that should be tested
type ResponseSample struct {
	Route        string
	ResponseCode int
}

type serveHTTP interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// ResponseCode tests the response code if it matches the expected ResponseCode of a sample for a given route
func ResponseCode(t *testing.T, r serveHTTP, samples []ResponseSample) {

	for _, sample := range samples {
		req, _ := http.NewRequest("GET", sample.Route, nil)
		recorder := httptest.NewRecorder()

		r.ServeHTTP(recorder, req)

		recorder.Flush()

		if recorder.Code != sample.ResponseCode {
			t.Errorf("response code %d on route %v, expected: %d", recorder.Code, sample.Route, sample.ResponseCode)
		}
	}
}
