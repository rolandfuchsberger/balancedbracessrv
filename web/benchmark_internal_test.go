package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type reqHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

func BenchmarkHTMLTemplate(b *testing.B) {

	tr, _ := newTemplateRenderer()

	benchmarkRenderer(tr, b)
}

func BenchmarkHTMLTemplateFast(b *testing.B) {

	tr, _ := newTemplateRendererFast()

	benchmarkRenderer(tr, b)
}

func BenchmarkPlushTemplate(b *testing.B) {

	tr, _ := newTemplateRendererPlush()

	benchmarkRenderer(tr, b)
}

// from fib_test.go
func benchmarkRenderer(handler reqHandler, b *testing.B) {

	// build request
	r, _ := http.NewRequest("GET", "/?expression=asdf", nil)

	recorder := httptest.NewRecorder()

	// fire request b.N times
	for n := 0; n < b.N; n++ {

		handler.Handle(recorder, r)

		//flush recorder
		recorder.Flush()
	}
}
