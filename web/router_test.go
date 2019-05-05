package web_test

import (
	"testing"

	"fuchsberger.email/balancedbracessrv/test"
	"fuchsberger.email/balancedbracessrv/web"
)

func TestRouterResponseCode(t *testing.T) {

	samples := []test.ResponseSample{
		{Route: "/html", ResponseCode: 200},
		{Route: "/fast", ResponseCode: 200},
		{Route: "/plush", ResponseCode: 200},
		{Route: "/wrongRoute", ResponseCode: 404},
	}

	test.ResponseCode(t, web.Routes(), samples)

}
