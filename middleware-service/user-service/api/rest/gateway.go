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
	router.Mount("/api/users", restRoutes(rest))
	router.Mount("/api/auth", authRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.Authenticator)
	r.Get("/{username}", rest.GetOne)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Post("/", rest.Post)
	r.Put("/", rest.Put)
	r.Patch("/", rest.Patch)
	r.Delete("/{userId}", rest.Delete)
	return
}

func authRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Post("/sign-up", rest.SingUp)
	r.Post("/login", rest.Login)
	r.Get("/refresh", rest.RefreshToken)
	r.Get("/logout", rest.Logout)
	return
}
