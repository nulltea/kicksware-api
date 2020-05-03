package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest RestfulHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)
	router.Mount("/api/references/sneakers", restRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/{referenceId}", rest.GetOne)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Post("/", rest.Post)
	r.Post("/multiply", rest.Post)
	r.Patch("/", rest.Patch)
	r.Get("/count", rest.Count)
	r.Post("/count", rest.Count)
	return
}
