package etc

import (
	"database/sql"
	"fmt"
)

var UseDB string

const ErrNameConlict = "Conflict"

var DB1 *sql.DB

func InitDBEtc1(dbs string) {
	var err error
	DB1, err = sql.Open("postgres", dbs)
	if err != nil {
		fmt.Println(`StaticTest - err`)
	}
	//	defer DB1.Close()

}
