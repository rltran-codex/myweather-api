package server

import (
	"context"
	"net/http"
	"time"

	"github.com/rltran-codex/myweather-api/internal/api/rest/routes"
	"github.com/rltran-codex/myweather-api/internal/cache"
	c "github.com/rltran-codex/myweather-api/internal/config"
)

var ctx = context.Background()

func Run() {
	// init redis client with 60 minute ttl cache
	cache.Connect(ctx, 60)

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         c.Config.Address,
		Handler:      mux,
		ReadTimeout:  time.Duration(c.Config.ReadTimeout * int(time.Second)),
		WriteTimeout: time.Duration(c.Config.WriteTimeout * int(time.Second)),
	}

	mux.HandleFunc("/api/v1/weather", routes.RateLimit(routes.GetWeatherByCity))
	routes.InitRate()
	routes.InitWeatherAPI()

	server.ListenAndServe()
}
