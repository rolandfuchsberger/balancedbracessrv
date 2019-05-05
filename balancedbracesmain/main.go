package balancedbracesmain

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"fuchsberger.email/balancedbracessrv/api"
	"fuchsberger.email/balancedbracessrv/web"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		//middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer, // Recover from panics without crashing server
	)

	router.Route("/api/", func(r chi.Router) {
		r.Mount("/balancedbraces/v1", api.Routes())
	})

	router.Route("/web/", func(r chi.Router) {
		r.Mount("/", web.Routes())
	})

	router.Get("/", handleRoot)

	return router
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	http.Redirect(w, req, "/web/html", http.StatusTemporaryRedirect)
}

// Main entry point - starts a server at port 8080
func Main() {

	router := routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	srv := &http.Server{Addr: ":8080", Handler: router}
	defer srv.Close()

	log.Fatal(srv.ListenAndServe())

}
