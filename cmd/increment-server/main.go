package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/bappush/go-course-project/internal/config"
	"github.com/bappush/go-course-project/internal/http-server/handlers/increment"
	"github.com/bappush/go-course-project/internal/models/counters"
	"github.com/bappush/go-course-project/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// set up logger
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	// init counter storage
	counterStorage := storage.NewCounterStorage()

	// init router
	router := chi.NewRouter()

	// handlers
	router.Post("/inc", increment.New(log, counterStorage))

	// start http server
	log.Info("starting http-server", slog.String("address", cfg.HTTPServerAddress), slog.String("env", cfg.Env))

	srv := &http.Server{
		Addr:    cfg.HTTPServerAddress,
		Handler: router,
	}

	// log counter value
	go logCounter(counterStorage, *log)

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", slog.String("err", err.Error()))
	}

	log.Error("server stopped")
}

// logCounter periodically logs counter value
func logCounter(provider increment.CounterProvider, logger slog.Logger) {
	tickPeriod := 5 * time.Second
	ticker := time.NewTicker(tickPeriod)

	keys := counters.GetAllowedKeys()

	for range ticker.C {
		for _, key := range keys {
			counter := provider.GetCounter(key)
			logger.Info("log counter", slog.String("key", key), slog.String("value", fmt.Sprintf("%v", counter)))
		}
	}
}
