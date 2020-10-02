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
	router.Mount("/orders", restRoutes(rest))
	router.Mount("/health", healthRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.auth.Authenticator)
	r.Use(rest.auth.Authorizer)
	r.Get("/{orderID}", rest.GetOne)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Post("/", rest.Post)
	r.Post("/multiply", rest.Post)
	r.Patch("/", rest.Patch)
	r.Get("/count", rest.Count)
	r.Post("/count", rest.Count)
	return
}

func healthRoutes(rest *Handler) (r *chi.Mux)  {
	r = chi.NewRouter()
	r.Get("/live", rest.HealthZ)
	r.Get("/ready", rest.ReadyZ)
	return
}