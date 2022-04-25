package server

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/handlers"
	"net/http"
)

func MainServer() {
	http.HandleFunc("/", handlers.MainHandlFunc)
	http.ListenAndServe(":"+config.Port, nil)
}
