package handlers

import (
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	_ "github.com/lib/pq"
	"net/http"
)

func PingGet(w http.ResponseWriter, r *http.Request) {

	//	connStr := "user=postgres password=qwerty dbname=goyp sslmode=disable"
	//	fmt.Println("DB string" + config.DBStringConnect)

	db, err := sql.Open("postgres", config.DBStringConnect)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	w.WriteHeader(http.StatusOK)
}
