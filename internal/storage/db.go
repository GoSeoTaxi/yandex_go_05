package storage

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
)

var StringConnect string

type Storage interface {
	PutDB(user, data string) (int, error)
	GetDB(idItem int) (string, error)
	GetDBLogin(login string) map[int]string
	ResoreDB(file string) error
	CheckLoginDB(login string) string
}

type countID struct {
	count int
}

//type DataPut struct {
//	URL1 string
//}

//type DataGet struct {
//	IDURLRedirect int
//}

//Хранение значений
var bd = map[int]string{}
var useBD = map[string]map[int]string{}

//var useBD = map[string]map[int]string{}
var index int
var fileNameDB string

func ResoreDB(fileName string) (status string, err error) {
	if len(fileName) < 1 {
		status = "NO FILE"
		return status, err
	}
	fmt.Println(`ПРоверка что файл передаётся _ ` + fileName)
	fileNameDB = fileName
	_, err = os.Stat(fileName)
	if err != nil {
		err := os.WriteFile(fileName, []byte("FILE_DB \n"), 0644)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
		status = "CREATED NEW FILE"
		return status, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		status = "err Open File"
		return status, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(`++++++++++++`)
	fmt.Println(`Читаю файл`)
	fmt.Println(`++++++++++++`)

	for scanner.Scan() {
		inputMap := strings.Split(scanner.Text(), "|http")
		if len(inputMap) == 2 {
			intKEYDB, _ := strconv.Atoi(inputMap[0])
			bd[intKEYDB] = "http" + inputMap[1]
		}
	}
	file.Close()
	return status, err
}

func writeFile(indInt int, data string) {
	f, _ := os.OpenFile(fileNameDB, os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	f.WriteString(strconv.Itoa(indInt) + "|" + data + "\n")
}

func PutDB(login, str string) (out int, err error) {

	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(`err sql open`)
	}
	defer db.Close()
	ping := db.Ping()
	if len(StringConnect) > 1 && err == nil && ping == nil {

		rowIDCount := db.QueryRow("select COUNT(id) from shortyp10")
		prodID := countID{}
		err = rowIDCount.Scan(&prodID.count)
		if err != nil {
			index = len(bd)
		} else {
			_, err := db.Exec("insert into shortyp10 (link, login) values ($1, $2)",
				str, login)
			if err != nil {
				fmt.Println(`err put`)
				index = len(bd)
			} else {
				//		fmt.Println(`use PUT db`)
				index = prodID.count
				//		fmt.Println(index)
			}
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
	//Тут добавить условие, если переменная окружения с файлом больше 1, то записываем данные в файл через функцию

	return index, err
}

func PutDBUni(login, str string) (out int, err error) {
	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(`err sql open`)
	}
	defer db.Close()
	ping := db.Ping()
	if len(StringConnect) > 1 && err == nil && ping == nil {

		outInt, err := PutPQ(str, login, StringConnect)
		if err != nil {
			if err.Error() == "Conflict" {

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

	return index, err
}

func GetDB(id int) (url2Redirect string, err error) {

	if etc.UseDB == "Y" {

		var linkVar string
		var isDelVar bool

		fmt.Println(`+++++USE DB+++++++`)

		db, err := sql.Open("postgres", StringConnect)
		if err != nil {
			fmt.Println(`err sql open`)
		}
		defer db.Close()

		fmt.Println(`+++++USE ID`)
		fmt.Println(id)

		err = db.QueryRow("SELECT link, is_del FROM shortyp10 where id = $1", id).Scan(&linkVar, &isDelVar)
		if err != nil {
			fmt.Println(err)
			url2Redirect = bd[id]
		} else {
			url2Redirect = linkVar

			if isDelVar != false {
				url2Redirect = ""
				err = fmt.Errorf("410")
				return url2Redirect, err
			}
			return url2Redirect, err
		}

	} else {
		url2Redirect = bd[id]
	}
	return url2Redirect, err
}

func GetDBLogin(loginInt string) (mapURL map[int]string) {
	mapURL = useBD[loginInt]
	//	fmt.Println(map1)
	return mapURL
}

func CheckLoginDB(login string) (status string) {

	//	fmt.Println(login)
	map1 := useBD[login]
	//	fmt.Println(len(map1))
	if len(map1) > 0 {
		status = "Y"
	}
	//	fmt.Println(status)
	return status
}
