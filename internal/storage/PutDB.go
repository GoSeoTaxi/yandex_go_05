package storage

import (
	"database/sql"
	"fmt"
)

type PutDBer interface {
	PutDBs() (int, error)
}

type PutDBT struct {
	LoginCookie string
	Links       string
}

func (p PutDBT) PutDBs() (out int, err error) {
	//func PutDB(login, str string) (out int, err error) {

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
	rowIDCount := db.QueryRow("select max(id) from shortyp10")
	prodID := countID{}
	err = rowIDCount.Scan(&prodID.count)

	fmt.Println(`++++++++!`)
	fmt.Println(`Проверка строк в бд`)
	fmt.Println(prodID.count)
	fmt.Println(`Печать err`)
	fmt.Println(err)
	fmt.Println(`++++++++!`)

	if err != nil {
		index = len(bd)
	} else {

		var ans int

		err = db.QueryRow("INSERT INTO shortyp10 (link, login) values ($1, $2)  RETURNING id", str, login).Scan(&ans)

		if err != nil {
			fmt.Println(`err put`)
			index = len(bd)
		} else {
			index = ans
			//		fmt.Println(index)
		}
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
	//Тут добавить условие, если переменная окружения с файлом больше 1, то записываем данные в файл через функцию

	return index, err
}
