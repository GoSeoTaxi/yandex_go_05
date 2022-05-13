package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
	"os"
)

func main() {
	SERVER_ADDRESS := os.Getenv("SERVER_ADDRESS")
	BASE_URL := os.Getenv("BASE_URL")

	config.LoadConfig(SERVER_ADDRESS, BASE_URL)
	server.MainServer()
}
