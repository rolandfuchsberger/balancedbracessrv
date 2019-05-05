package balancedbracesmain_test

import (
	"net/http"
	"testing"

	"fuchsberger.email/balancedbracessrv/balancedbracesmain"
)

func TestMainStartsServer(t *testing.T) {

	// start server
	go balancedbracesmain.Main()

	// test http request on :8080
	url := "http://localhost:8080/"
	_, err := http.Get(url)
	if err != nil {
		t.Errorf("Error sending get request to %v: %v", url, err)
	}

}
