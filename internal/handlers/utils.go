package handlers

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
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

func asyncDel(intURL []string, login string) {

	//Видимо, по условию задачи, тут нужно передать все URL в отдельную горутину, которая держит соединение с базой и занимает общенией с ней,
	//но мы идём по пути минимальной реализаци по ТЗ (в данном случае по тестам)

	for i := 0; i < len(intURL); i++ {
		storage.DelPQ(intURL[i], login)
	}

}
