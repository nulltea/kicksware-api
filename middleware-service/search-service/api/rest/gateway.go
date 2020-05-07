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
	router.Mount("/api/search/reference", restRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
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
