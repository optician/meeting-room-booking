package internal

import (
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	"github.com/optician/meeting-room-booking/internal/administration/db"
	"github.com/optician/meeting-room-booking/internal/administration/httpapi"
	"github.com/optician/meeting-room-booking/internal/administration/service"
	"github.com/optician/meeting-room-booking/internal/dbPool"
	"go.uber.org/zap"
)

func Make(httpLogger *httplog.Logger, logger *zap.SugaredLogger, config *Config) chi.Router {
	dbPool, dbPoolErr := dbPool.NewDBPool(&config.DB, logger)
	if dbPoolErr != nil {
		logger.Fatalf("application terminated: %v", dbPoolErr)
		os.Exit(-1)
	}
	// defer dbPool.Close()  // how to close correctly?

	roomsDB := db.New(dbPool.GetPool(), logger)
	adminLogic := service.Make(&roomsDB, logger)

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
		httplog.RequestLogger(httpLogger),
		middleware.CleanPath,
		middleware.ContentCharset(allowedCharsets...),
		cors.Handler(corsOptions),
		middleware.Timeout(20*time.Second),
		middleware.Recoverer,
	)

	r.Group(httpapi.Make(&adminLogic, logger))

	return r
}
