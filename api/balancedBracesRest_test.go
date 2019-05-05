package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"fuchsberger.email/balancedbracessrv/api"
	"fuchsberger.email/balancedbracessrv/test"
)

// TestRouterResponseCode tests if the ResponseCode for a given route matches the expectation
func TestRouterResponseCode(t *testing.T) {

	samples := []test.ResponseSample{
		{Route: "/asdf", ResponseCode: 200},
		{Route: "/123456", ResponseCode: 200},
		{Route: "/plush", ResponseCode: 200},
		{Route: "/", ResponseCode: 200},
		{Route: "/wrongRoute/asdf", ResponseCode: 404},
	}

	test.ResponseCode(t, api.Routes(), samples)
}

// TestResponse tests if the rest api responds to a given expression with the correct json
func TestResponse(t *testing.T) {

	samples := []struct {
		expr           string
		balancedbraces bool
	}{
		{"asdf", true},
		{")(}", false},
		{"a(sdf", false},
		{"as(d]f", false},
		{"([])", true},
		{"({[asdf]})", true},
		{"([asdf/]})", false},
		{"({[asdf]}", false},
		{"({[[{(asdf", false},
		{"(", false},
		{"", true},
		{"(/){}[]", true},
		{"(){}[", false},
		{"(){[]", false},
		{"({}[]", false},
	}

	r := api.Routes()

	for _, sample := range samples {

		req, _ := http.NewRequest("GET", "/"+url.PathEscape(sample.expr), nil)
		recorder := httptest.NewRecorder()

		r.ServeHTTP(recorder, req)

		recorder.Flush()

		// pointer for the decoder to decode the response body
		response := &struct {
			BalancedBraces bool `json:"balancedBraces"`
		}{}

		json.NewDecoder(recorder.Body).Decode(response)

		if response.BalancedBraces != sample.balancedbraces {
			t.Errorf("Request URL: %v, Expected Response: %v, Received: %v", "/"+sample.expr, sample.balancedbraces, response.BalancedBraces)
		}

	}

}
