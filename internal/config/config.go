package config

import (
	"fmt"
	"strings"
)

var ServerHost string
var Port string
var ConstGetEndPoint string
var PathURLConf string

func LoadConfig(adr, path string) {

	if len(adr) < 2 {
		Port = "8080"
	} else {
		strAdr := strings.Split(adr, ":")
		Port = strAdr[1]
	}

	if len(path) < 16 {
		ServerHost = "http://localhost"
	} else {
		strPath := strings.Split(path, ":")
		Port = strPath[2]
		ServerHost = strPath[0] + ":" + strPath[1]
	}
	PathURLConf = "/"
	ConstGetEndPoint = "id"

	fmt.Println(`I am Starting ` + `server ` + ServerHost + ` port - ` + Port + ` path - ` + PathURLConf + ` endPoint - ?` + ConstGetEndPoint + `=`)

}
