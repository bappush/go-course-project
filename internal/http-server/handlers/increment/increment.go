package increment

import (
	"github.com/bappush/go-course-project/internal/lib/response"
	"github.com/bappush/go-course-project/internal/models/counters"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Response struct {
	response.Response
	Counter int `json:"counter"`
}

type CounterProvider interface {
	Increment(key string)
	GetCounter(key string) int
	GetCounters() map[string]int
}

// New creates a new HandlerFunc for POST /inc
func New(log *slog.Logger, provider CounterProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counterKey := r.URL.Query().Get("name")
		if !counters.IsKeyAllowed(counterKey) {
			log.Error("incorrect counter key", slog.String("key", counterKey))

			render.JSON(w, r, Response{
				Response: response.Error("Incorrect counter key"),
			})

			return
		}

		provider.Increment(counterKey)

		updatedCounter := provider.GetCounter(counterKey)

		render.JSON(w, r, Response{
			Response: response.OK(),
			Counter:  updatedCounter,
		})
	}
}
