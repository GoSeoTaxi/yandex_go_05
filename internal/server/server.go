package server

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"net/http"
)

func MainServer() {

	// Logger
	logger := httplog.NewLogger("httplog", httplog.Options{
		LogLevel: "Debug",
		JSON:     true,
		Concise:  true,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
	})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Use(httplog.RequestLogger(logger))

	r.Use(middleware.Compress(1, "gzip"))

	r.Get(config.PathURLConf, handlers.MainHandlFuncGet)
	r.Post(config.PathURLConf, handlers.MainHandlFuncPost)
	r.Post("/api/shorten", handlers.APIJSON)
	http.ListenAndServe(":"+config.Port, r)

}
