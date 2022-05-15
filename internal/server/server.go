package server

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func MainServer() {
	r := chi.NewRouter()
	r.Use(middleware.Compress(5))
	r.Get(config.PathURLConf, handlers.MainHandlFunc)
	r.Post("/", handlers.MainHandlFunc)
	r.Post("/api/shorten", handlers.APIJSON)
	http.ListenAndServe(":"+config.Port, r)

}
