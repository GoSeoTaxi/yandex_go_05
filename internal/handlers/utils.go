package handlers

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
)

func MakeString(idItem string) string {
	return config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idItem
}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(`err - create Token`)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
