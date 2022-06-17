package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
)

func main() {
	config.InitCLI()
	server.MainServer()

	/*

		Из того, что я понял из "основных" проблем

		Интерфейс - не интерфейс, а набор функций

		Нет пакета запросов
		БД создаёт новые коннекты

	*/
}
