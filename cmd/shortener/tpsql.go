package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(`est`)

	connStr := "user=postgres password=qwerty dbname=goyp1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	/*
		result, err := db.Exec("insert into shortyp10 (link) values ($1)",
			"https://yhu.cym")
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected()) // количество добавленных строк
	*/
}
