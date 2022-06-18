package storage

import (
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
)

type GetDBer interface {
	GetDBs() (string, error)
}

type GetDBT struct {
	IDItem int
}

func (p GetDBT) GetDBs() (url2Redirect string, err error) {

	id := p.IDItem

	if etc.UseDB == "Y" {

		var linkVar string
		var isDelVar bool

		db, err := sql.Open("postgres", StringConnect)
		if err != nil {
			fmt.Println(`err sql open`)
		}
		defer db.Close()

		err = db.QueryRow("SELECT link, is_del FROM shortyp10 where id = $1", id).Scan(&linkVar, &isDelVar)
		if err != nil {
			fmt.Println(err)
			url2Redirect = bd[id]
			return url2Redirect, err
		}

		url2Redirect = linkVar
		if isDelVar != false {
			url2Redirect = ""
			err = fmt.Errorf("410")
			return url2Redirect, err
		}
		return url2Redirect, err

	} else {
		url2Redirect = bd[id]
	}
	return url2Redirect, err
}
