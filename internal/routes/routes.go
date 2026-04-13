package routes

import (
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/handlers"
	"github.com/Thomazoide/IACam_backend/internal/ws"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/cameras", func(r chi.Router) {
		r.Post("/create", handlers.CreateCamera)
		r.Get("/", handlers.GetCameras)
		r.Get("/{id}", handlers.GetCamera)
		r.Delete("/delete/{id}", handlers.DeleteCamera)
	})

	r.Get("/stream/{id}", handlers.StreamProxy)
	r.Post("/events", handlers.CatchEvent)
	r.Get("/ws", ws.HandleWS)

	return r
}
