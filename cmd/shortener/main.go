package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
)

func main() {

	SERVER_ADDRESS := "127.0.0.1:8081"
	BASE_URL := "/base/"

	config.LoadConfig(SERVER_ADDRESS, BASE_URL)
	server.MainServer()
}
