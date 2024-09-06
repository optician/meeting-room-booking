package httpapi

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	"github.com/optician/meeting-room-booking/internal/administration/httpapi"
)

func Make(logger *httplog.Logger) chi.Router {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	allowedCharsets := []string{"UTF-8", "Latin-1", ""}

	r.Use(
		middleware.Heartbeat("/liveness"),
		middleware.Heartbeat("/readiness"), // usually it's ok, especially because I don't have a k8s at the moment
		httplog.RequestLogger(logger),
		middleware.CleanPath,
		middleware.ContentCharset(allowedCharsets...),
		cors.Handler(corsOptions),
		middleware.Timeout(20*time.Second),
		middleware.Recoverer,
	)

	r.Group(httpapi.Make)

	return r
}
