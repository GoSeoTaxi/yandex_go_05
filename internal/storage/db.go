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

//Функция дозаписи в файл
func writeFile(indInt int, data string) {
	f, _ := os.OpenFile(fileNameDB, os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	f.WriteString(strconv.Itoa(indInt) + "|" + data + "\n")
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
