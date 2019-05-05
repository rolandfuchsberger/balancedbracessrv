package web

import (
	"log"

	"github.com/go-chi/chi"
)

//Routes asdf
func Routes() *chi.Mux {
	router := chi.NewRouter()

	{
		bbhandlerPlush, err := NewTemplateRendererPlush()
		if err != nil {
			log.Panicln(err)
		}
		router.Get("/plush", bbhandlerPlush.Handle)
	}

	{
		bbhandler2, err := NewTemplateRendererFast()
		if err != nil {
			log.Panicln(err)
		}
		router.Get("/fast", bbhandler2.Handle)
	}

	{
		bbhandler, err := NewTemplateRenderer()
		if err != nil {
			log.Panicln(err)
		}

		router.Get("/html", bbhandler.Handle)
	}

	return router
}
