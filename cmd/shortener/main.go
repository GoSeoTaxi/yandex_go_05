package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
	"os"
)

func main() {
	config.LoadConfig(os.Getenv("SERVER_ADDRESS"), os.Getenv("BASE_URL"))
	server.MainServer()
}
