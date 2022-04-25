package handlers

import "github.com/GoSeoTaxi/yandex_go_05/internal/config"

func MakeString(idItem string) string {
	return config.Server_host + ":" + config.Port + "/?" + config.ConstGetEndPoint + "=" + idItem
}
