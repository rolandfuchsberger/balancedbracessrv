package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fuchsberger.email/balancedbracessrv/web"
)

type reqHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

func BenchmarkHTMLTemplate(b *testing.B) {

	tr, _ := web.NewTemplateRenderer()

	benchmarkRenderer(tr, b)
}

func BenchmarkHTMLTemplateFast(b *testing.B) {

	tr, _ := web.NewTemplateRendererFast()

	benchmarkRenderer(tr, b)
}

func BenchmarkPlushTemplate(b *testing.B) {

	tr, _ := web.NewTemplateRendererPlush()

	benchmarkRenderer(tr, b)
}

// from fib_test.go
func benchmarkRenderer(handler reqHandler, b *testing.B) {

	r, _ := http.NewRequest("GET", "/?expression=asdf", nil)

	//r.URL.Query().Add("expression", "asdf")

	recorder := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {

		handler.Handle(recorder, r)

		//flush recorder
		recorder.Flush()
	}
}
