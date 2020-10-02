package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest *Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)
	router.Mount("/search/reference", restRoutes(rest))
	router.Mount("/health", healthRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.auth.Authenticator)
	r.Get("/", rest.Get)
	r.Get("/by/{field}", rest.GetBy)
	r.Get("/sku/{sku}", rest.GetSKU)
	r.Get("/brand/{brand}", rest.GetBrand)
	r.Get("/model/{model}", rest.GetModel)
	r.Post("/{referenceId}", rest.PostOne)
	r.Post("/", rest.Post)
	r.Post("/all", rest.PostAll)
	r.Post("/query", rest.PostQuery)
	return
}

func healthRoutes(rest *Handler) (r *chi.Mux)  {
	r = chi.NewRouter()
	r.Get("/live", rest.HealthZ)
	r.Get("/ready", rest.ReadyZ)
	return
}