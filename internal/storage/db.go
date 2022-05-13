package storage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Storage interface {
	PutDB(data string) (int, error)
	GetDB(idItem int) (string, error)
	ResoreDB(file string) error
}

//type DataPut struct {
//	URL1 string
//}

//type DataGet struct {
//	IDURLRedirect int
//}

//Хранение значений
var bd = map[int]string{}
var index int
var fileNameDB string

func ResoreDB(fileName string) (status string, err error) {
	fmt.Println(fileName + `ПРоверка что файл передаётся`)
	if len(fileName) < 1 {
		status = "NO FILE"
		return status, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println(`ERR - create file`)
			panic(err)
		}
		return "newCreate", err
		defer f.Close()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileNameDB = fileName

	for scanner.Scan() {
		inputMap := strings.Split(scanner.Text(), "|http")
		intKEYDB, _ := strconv.Atoi(inputMap[0])
		bd[intKEYDB] = "http" + inputMap[1]
	}

	if err := scanner.Err(); err != nil {
		return status, err
	}

	return status, err
}

//Функция дозаписи в файл
func writeFile(indInt int, data string) {
	f, _ := os.OpenFile(fileNameDB, os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString("\n" + strconv.Itoa(indInt) + "|" + data)
	defer f.Close()
}

func PutDB(str string) (out int, err error) {
	index = len(bd)
	bd[index] = str
	if fileNameDB != "" {
		writeFile(index, str)
	}
	//Тут добавить условие, если переменная окружения с файлом больше 1, то записываем данные в файл через функцию

	return index, err
}

func GetDB(id int) (url2Redirect string, err error) {
	return bd[id], err
}
