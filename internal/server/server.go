package server

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func MainServer() {

	r := chi.NewRouter()
	r.Get("/", handlers.MainHandlFunc)
	r.Post("/", handlers.MainHandlFunc)
	r.Post("/api/shorten/", handlers.ApiJson)
	http.ListenAndServe(":"+config.Port, r)

}
