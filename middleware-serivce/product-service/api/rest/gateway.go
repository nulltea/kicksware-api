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
	router.Mount("/products/sneakers", restRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/{sneakerId}", rest.GetOne)
	r.Get("/query", rest.Get)
	r.Get("/map", rest.Get)
	r.Get("/", rest.GetAll)
	r.Post("/", rest.Post)
	r.Put("/", rest.Put)
	r.Patch("/", rest.Patch)
	r.Delete("/{sneakerId}", rest.Delete)
	return
}
