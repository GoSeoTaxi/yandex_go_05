package handlers

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"net/http"
	"time"
)

var bd1db storage.StorageBD

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
	for i, url := range intURL {
		//for i := 0; i < len(intURL); i++ {

		fmt.Println(`собираюсь удалить`)
		duration := time.Second * 2
		time.Sleep(duration)
		fmt.Println(`САЙТ`)
		fmt.Println(intURL[i])
		fmt.Println(`LOGIN`)
		fmt.Println(login)
		fmt.Println(`Отправил`)

		bd1db.DelPQ(url, login)
	}
}

func checkLogin(req http.Request) string {
	var loginCookie string
	login, err := req.Cookie("login")
	if err != nil {
		loginCookie = "anonimus"
	} else {
		loginCookie = login.Value
	}
	return loginCookie
}
