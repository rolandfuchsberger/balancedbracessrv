package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"fuchsberger.email/balancedbracessrv/balancedbraces"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Routes returns a chi router that handles requests to the balanced braces api
// adds also a CORS header
func Routes() *chi.Mux {
	router := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router.Use(
		cors.Handler,
		middleware.Logger, // Log API request calls
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		//middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
	)

	router.Get("/{expr}", balancedBracesRest)

	// Handle also if expr = ""
	router.Get("/", handleRoot)

	return router
}

// handleRoot handles requests with empty expressions and forwards them to balancedBracesRest
func handleRoot(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	// forward to balancedBraces
	balancedBracesRest(w, req)
}

// balancedBracesRest extracts the expression from a request and returns a json containing information if the expression has balanced braces
func balancedBracesRest(w http.ResponseWriter, r *http.Request) {

	expr, err := url.QueryUnescape(chi.URLParam(r, "expr"))

	if err != nil {
		log.Printf("Error in api.balancedBracesRest: %v", err)
		return
	}

	ret := struct {
		BalancedBraces bool `json:"balancedBraces"`
	}{
		balancedbraces.BalancedBraces(expr),
	}

	json.NewEncoder(w).Encode(ret)

}
