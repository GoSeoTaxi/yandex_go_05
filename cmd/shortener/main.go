package main

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/server"
)

func main() {
	config.InitCLI()
	server.MainServer()

	/*

				Исправленные проблемы -
				Интерфейс - не интерфейс, а набор функций storage.GetDBLogin - готово
				БД создаёт новые коннекты
				Нет пакета запросов

				Обсудить
		1) Правильность использования постоянного подключения к бд /etc/dbps.go
		2) Ошибки в статическом тесте - composite literal uses unkeyed fields
		3) Указать что почитать про continue
		4) Обсудить отличие
			for i := 0; i < len(intURL); i++ {
			от for i, url := range intURL {
		5) Разобраться с тестами, почему приходит пустой запрос
		6) Как запускать тесты локально? (не на GitHub)

	*/
}
