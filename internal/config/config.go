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

	if len(adr) < 9 {
		ServerHost = "http://" + "127.0.0.1"
		Port = "8080"
	} else {

		str := strings.Split(adr, ":")
		ServerHost = "http://" + str[0]
		Port = str[1]

	}

	if len(path) < 1 {
		PathURLConf = "/"
	} else {
		PathURLConf = path
	}

	ConstGetEndPoint = "id"

	fmt.Println(`I am Starting ` + `server ` + ServerHost + ` port - ` + Port + ` path - ` + PathURLConf + ` endPoint - ?` + ConstGetEndPoint + `=`)

}
