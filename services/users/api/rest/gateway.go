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
	router.Mount("/users", restRoutes(rest))
	router.Mount("/auth", authRoutes(rest))
	router.Mount("/mail", mailRoutes(rest))
	router.Mount("/interact", interactRoutes(rest))
	router.Mount("/health", healthRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.Authenticator)
	r.Use(rest.Authorizer)
	r.Get("/{userID}", rest.GetOne)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Post("/", rest.Post)
	r.Put("/", rest.Put)
	r.Patch("/", rest.Patch)
	r.Delete("/{userID}", rest.Delete)
	return
}

func authRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Post("/sign-up", rest.SingUp)
	r.Post("/login", rest.Login)
	r.Post("/remote", rest.Remote)
	r.Get("/guest", rest.Guest)
	r.Get("/token-refresh", rest.RefreshToken)
	r.Get("/logout", rest.Logout)
	return
}

func mailRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.Authenticator)
	r.Use(rest.Authorizer)
	r.Get("/confirm", rest.SendEmailConfirmation)
	r.Get("/password-reset", rest.SendResetPassword)
	r.With(rest.Authorizer).Get("/notify", rest.SendNotification)
	return
}

func interactRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.Authenticator)
	r.Use(rest.Authorizer)
	r.Get("/like/{entityID}", rest.Like)
	r.Get("/unlike/{entityID}", rest.Unlike)
	return
}

func healthRoutes(rest *Handler) (r *chi.Mux)  {
	r = chi.NewRouter()
	r.Get("/live", rest.HealthZ)
	r.Get("/ready", rest.ReadyZ)
	return
}