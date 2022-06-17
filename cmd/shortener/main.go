package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
)

func main() {
	config.InitCLI()
	server.MainServer()

	/*

		У тебя здесь не происходит использования интерфейса хранилища - нет объекта. И в итоге прямые зависимости на конкретные реализации - Нужно это обсудить другими словами.

	*/
}
