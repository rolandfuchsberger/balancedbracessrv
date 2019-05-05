package web

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandleErrorInTemplate(handler reqHandler, t *testing.T) {

	r, _ := http.NewRequest("GET", "/", nil)

	recorder := httptest.NewRecorder()

	var logOut bytes.Buffer
	log.SetOutput(&logOut)

	handler.Handle(recorder, r)

	//close recorder
	recorder.Flush()

	if logOut.Bytes() == nil || len(logOut.Bytes()) == 0 {
		t.Error("No error message received on stdout")
	}

	if recorder.Code != 500 {
		t.Errorf("response code: %v, expected: 500", recorder.Code)
	}
}

func testHandle(handler reqHandler, t *testing.T) {

	var response1Body, response2Body string

	//#### Test response without expression
	{
		r, _ := http.NewRequest("GET", "/", nil)

		recorder := httptest.NewRecorder()

		handler.Handle(recorder, r)

		//close recorder
		recorder.Flush()

		response1Body = recorder.Body.String()
	}

	//#### Test response with expression
	{
		r, _ := http.NewRequest("GET", "/?expression=asdf", nil)

		//r.URL.Query().Add("expression", "asdf")

		recorder := httptest.NewRecorder()

		handler.Handle(recorder, r)

		//close recorder
		recorder.Flush()

		response2Body = recorder.Body.String()
	}

	if response1Body == response2Body {
		t.Errorf("expected different response bodys for different expressions. Response1: %q Response2: %q",
			response1Body, response2Body)
	}
}
