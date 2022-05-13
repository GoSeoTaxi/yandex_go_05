package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
)

func main() {
	SERVER_ADDRESS := "127.0.0.1:8080"
	BASE_URL := "/"

	config.LoadConfig(SERVER_ADDRESS, BASE_URL)
	server.MainServer()
}
