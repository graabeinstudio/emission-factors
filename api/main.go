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

type EmissionFactorTypesOutput struct {
	Body struct {
		EmissionFactors []internal.EmissionFactorType `json:"emissionFactorsTypes" doc:"Available emission factors types"`
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

	cli := huma.NewCLI(func(hooks huma.Hooks, options *Options) {
		router := chi.NewMux()
		router.Use(httplog.RequestLogger(logger))

		config := huma.DefaultConfig("Norwegian emission factors", appVersion)
		config.DocsPath = "/"
		api := humachi.New(router, config)
		

		huma.Get(api, "/emission-factors", func(ctx context.Context, input *struct{
			Name string `doc:"List all available emission factors"`
		}) (*EmissionFactorsOutput, error) {
			resp := &EmissionFactorsOutput{}
			resp.Body.EmissionFactors = internal.EmissionFactorsNorway2022
			return resp, nil
		})

		huma.Get(api, "/emission-factor-types", func(ctx context.Context, input *struct{
			Name string `doc:"List all available emission factors"`
		}) (*EmissionFactorTypesOutput, error) {
			resp := &EmissionFactorTypesOutput{}
			resp.Body.EmissionFactors = internal.EmissionFactorTypes
			return resp, nil
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: router,
		}

		hooks.OnStart(func() {
			logger.Info(fmt.Sprintf("Starting server on port %d...", options.Port))
			server.ListenAndServe()
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	cli.Run()
}
