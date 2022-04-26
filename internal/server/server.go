package server

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func MainServer() {

	r := chi.NewRouter()
	r.Get("/", handlers.MainHandlFunc)
	r.Post("/", handlers.MainHandlFunc)

	http.ListenAndServe(":8080", r)

}
