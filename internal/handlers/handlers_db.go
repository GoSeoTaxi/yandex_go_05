package handlers

import (
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	_ "github.com/lib/pq"
	"net/http"
)

func PingGet(w http.ResponseWriter, r *http.Request) {

	//	connStr := "user=postgres password=qwerty dbname=goyp sslmode=disable"
	//	fmt.Println("DB string" + config.DBStringConnect)
	/*
		db, err := sql.Open("postgres", config.DBStringConnect)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer db.Close()
	*/

	c := storage.PingDBT{}
	resp := storage.StoragePing.PingDBTs(c)

	if resp != true {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
