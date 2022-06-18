package storage

import (
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
)

type PutDBUnier interface {
	PutDBUnis() (int, error)
}

type PutDBUniT struct {
	LoginCookie string
	Links       string
}

func (p PutDBUniT) PutDBUnis() (out int, err error) {

	login := p.LoginCookie
	str := p.Links

	fmt.Println(`++++++++++++++++++`)
	fmt.Println(login)
	fmt.Println(`++++++++++`)
	fmt.Println(str)
	fmt.Println(`++++++++++++++++++`)

	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(`err sql open`)
	}
	defer db.Close()
	ping := db.Ping()
	if len(StringConnect) > 1 && err == nil && ping == nil {
		outInt, err := PutPQ(str, login, StringConnect)
		if err != nil {
			if err.Error() == etc.ErrNameConlict {
				return outInt, err
			} else {
				fmt.Println(err)
				fmt.Println(`err - put dbpq`)
			}

		} else {
			index = outInt
		}
	} else {
		index = len(bd)
	}

	bd[index] = str

	map1 := useBD[login]
	if len(map1) == 0 {
		map1 = make(map[int]string)
	}
	map1[index] = str
	useBD[login] = map1

	if fileNameDB != "" {
		writeFile(index, str)
	}

	fmt.Println(`Что мы отдаём? - PutDBUni`)
	fmt.Println(index)

	return index, err
}
