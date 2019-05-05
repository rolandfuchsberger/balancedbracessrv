package web

import (
	"log"

	"github.com/go-chi/chi"
)

//Routes retunrs all routes handeled by this package
func Routes() *chi.Mux {
	router := chi.NewRouter()

	{
		bbhandlerPlush, err := newTemplateRendererPlush()
		if err != nil {
			log.Panicln(err)
		}
		router.Get("/plush", bbhandlerPlush.Handle)
	}

	{
		bbhandler2, err := newTemplateRendererFast()
		if err != nil {
			log.Panicln(err)
		}
		router.Get("/fast", bbhandler2.Handle)
	}

	{
		bbhandler, err := newTemplateRenderer()
		if err != nil {
			log.Panicln(err)
		}

		router.Get("/html", bbhandler.Handle)
	}

	return router
}
