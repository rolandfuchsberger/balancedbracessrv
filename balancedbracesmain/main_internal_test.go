package balancedbracesmain

import (
	"net/http"

	"testing"
	"fuchsberger.email/balancedbracessrv/test"
)

// TestInitRouter tests if the initialization of the router + walking the routs yiels any error
func TestInitRouter(t *testing.T) {

	initRouter()
}

// TestRouterResponseCode tests if the ResponseCode for a given route matches the expectation
func TestRouterResponseCode(t *testing.T) {

	samples := []test.ResponseSample{
		{Route: "/web/html", ResponseCode: 200},
		{Route: "/api/balancedbraces/v1/asdf", ResponseCode: 200},
		{Route: "/", ResponseCode: http.StatusTemporaryRedirect}, 
		{Route: "/wrongRoute/asdf", ResponseCode: 404},
		{Route: "/wrongRoute", ResponseCode: 404},
	}

	test.ResponseCode(t, routes(), samples)
}