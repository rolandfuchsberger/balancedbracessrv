package test_test

import (
	"net/http"
	"testing"

	"fuchsberger.email/balancedbracessrv/test"
)

type testRouter struct{}

func (r testRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(200)

}

// TestRouterTest tests the router test suit
func TestRouterTest(t *testing.T) {

	r := testRouter{}
	samples := []test.ResponseSample{
		{Route: "asdf", ResponseCode: 200},
	}

	test.ResponseCode(t, r, samples)

}
