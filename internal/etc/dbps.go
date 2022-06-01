package etc

/*
import (
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	_ "github.com/lib/pq"
)

func WriteDB(url, login string) (int, error) {

	db, err := sql.Open("postgres", config.DBStringConnect)
	defer db.Close()
	ping := db.Ping()

	if len(config.DBStringConnect) > 1 && err == nil && ping == nil {

		rowIDCount := db.QueryRow("select COUNT(id) from shortyp10")
		prodID := product{}
		err = rowIDCount.Scan(&prodID.count)

		fmt.Println(`use PUT db`)

		//		index = prodID.count

		idWrite := 3
		return idWrite, err
	} else {
		idWrite := 0
		err = ping
		return idWrite, err
	}

}
*/
