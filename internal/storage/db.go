package storage

import (
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
)

type StorageBD interface {
	ResoreDBFile() (string, error)
}

type RestoreDBFile struct {
	FileName string
}

func (p RestoreDBFile) ResoreDBFile() (status string, err error) {

	fileName := p.FileName

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

	return status, err
}

type countID struct {
	count int
}

var StringConnect string

//Хранение значений
var bd = map[int]string{
	0: "t0",
	1: "t1",
}

var useBD = map[string]map[int]string{}

//var useBD = map[string]map[int]string{}
var index int
var fileNameDB string

func writeFile(indInt int, data string) {
	f, _ := os.OpenFile(fileNameDB, os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	f.WriteString(strconv.Itoa(indInt) + "|" + data + "\n")
}

//func GetDBLogin(loginInt string) (mapURL map[int]string) {
//	mapURL = useBD[loginInt]
//	//	fmt.Println(map1)
//	return mapURL
//}

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
