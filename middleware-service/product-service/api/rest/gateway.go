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
		rest.Authenticator,
	)
	router.Mount("/products/sneakers", restRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/{sneakerId}", rest.GetOne)
	r.Get("/query", rest.Get)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.With(rest.Authorizer).Post("/", rest.Post)
	r.With(rest.Authorizer).Put("/", rest.Put)
	r.With(rest.Authorizer).Put("/{sneakerId}/images", rest.PutImages)
	r.With(rest.Authorizer).Patch("/", rest.Patch)
	r.With(rest.Authorizer).Delete("/{sneakerId}", rest.Delete)
	return
}
