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

func Make(logger *httplog.Logger) chi.Router {
	zapLogger := zap.NewExample().Sugar()

	dbPool, dbPoolErr := dbPool.NewDBPool(zapLogger)
	if dbPoolErr != nil {
		zapLogger.Fatalf("application terminated: %v", dbPoolErr)
		os.Exit(-1)
	}
	// defer dbPool.Close()  // how to close correctly

	roomsDB := db.New(dbPool.GetPool(), zapLogger)
	adminLogic := service.Make(&roomsDB, zapLogger)

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

	r.Group(httpapi.Make(&adminLogic, zapLogger))

	return r
}
