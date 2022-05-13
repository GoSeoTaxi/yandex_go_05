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
	if len(fileName) < 1 {
		status = "NO FILE"
		return status, err
	}
	fmt.Println(`ПРоверка что файл передаётся _ ` + fileName)

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			fmt.Println(`ERR - create file`)
			panic(err)
		}
		return "newCreate", err
	}

	scanner := bufio.NewScanner(file)
	fileNameDB = fileName

	for scanner.Scan() {
		inputMap := strings.Split(scanner.Text(), "|http")
		intKEYDB, _ := strconv.Atoi(inputMap[0])
		bd[intKEYDB] = "http" + inputMap[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return status, err
	}

	return status, err
}

//Функция дозаписи в файл
func writeFile(indInt int, data string) {
	f, _ := os.OpenFile(fileNameDB, os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(strconv.Itoa(indInt) + "|" + data + "\n")
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
