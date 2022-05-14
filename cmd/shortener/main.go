package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"os"
)

func main() {
	config.LoadConfig(os.Getenv("SERVER_ADDRESS"), os.Getenv("BASE_URL"))
	storage.ResoreDB(os.Getenv("FILE_STORAGE_PATH"))
	server.MainServer()
}
