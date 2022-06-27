package storage

import (
	"database/sql"
	"fmt"
)

type StoragePing interface {
	PingDBTs() bool
}

type PingDBT struct {
	Ping string
}

func (p PingDBT) PingDBTs() bool {

	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}
