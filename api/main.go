package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/graabeinstudio/emission-factors/internal"
)

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type EmissionFactorsOutput struct {
	Body struct {
		EmissionFactors []internal.EmissionFactor `json:"emissionFactors" doc:"Emission factors"`
	}
}

func main() {
	appVersion := "0.0.0"

	// Logger
	logger := httplog.NewLogger("emission-factors-api", httplog.Options{
		JSON:             true,
		LogLevel:         slog.LevelInfo,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		Tags: map[string]string{
		  "version": appVersion,
		  "env":     "dev",
		},
		QuietDownRoutes: []string{},
		QuietDownPeriod: 10 * time.Second,
	  })

	  // Create a CLI app which takes a port option.
	cli := huma.NewCLI(func(hooks huma.Hooks, options *Options) {
		// Create a new router & API
		router := chi.NewMux()
		router.Use(httplog.RequestLogger(logger))

		config := huma.DefaultConfig("Norwegian emission factors", appVersion)
		config.DocsPath = "/"
		api := humachi.New(router, config)
		

		// Add the operation handler to the API.
		huma.Get(api, "/emission-factors", func(ctx context.Context, input *struct{
			Name string `doc:"List all available emission factors"`
		}) (*EmissionFactorsOutput, error) {
			resp := &EmissionFactorsOutput{}
			resp.Body.EmissionFactors = internal.EmissionFactorsNorway2022
			return resp, nil
		})

		// Create the HTTP server.
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: router,
		}

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			logger.Info(fmt.Sprintf("Starting server on port %d...", options.Port))
			server.ListenAndServe()
		})

		// Tell the CLI how to stop your server.
		hooks.OnStop(func() {
			// Give the server 5 seconds to gracefully shut down, then give up.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
