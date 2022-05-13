package config

var ServerHost string
var Port string
var ConstGetEndPoint string
var PathURLConf string

func LoadConfig(adr, path string) {

	if len(adr) < 9 {
		ServerHost = "http://127.0.0.1"
	} else {
		ServerHost = adr
	}

	if len(path) < 1 {
		PathURLConf = "/"
	} else {
		PathURLConf = path
	}

	Port = "8080"
	ConstGetEndPoint = "id"

}
