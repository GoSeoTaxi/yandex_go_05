package config

import (
	"flag"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"os"
)

func InitCLI() {

	var serverAddress string
	var baseURL string
	var fileStoragePath string
	var dbStringConnect string

	flag.StringVar(&dbStringConnect, "d", "user=postgres password=qwerty dbname=goyp1 sslmode=disable", "DATABASE_DSN")
	flag.StringVar(&serverAddress, "a", "localhost:8080", "SERVER_ADDRESS")
	flag.StringVar(&baseURL, "b", ":8080", "BASE_URL")
	flag.StringVar(&fileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	var serverAddressVar string
	var baseURLVar string
	var fileStoragePathVar string
	var dbStringConnectVar string

	serverAddressVar = os.Getenv("SERVER_ADDRESS")
	baseURLVar = os.Getenv("BASE_URL")
	fileStoragePathVar = os.Getenv("FILE_STORAGE_PATH")
	dbStringConnectVar = os.Getenv("DATABASE_DSN")

	if len(dbStringConnect) < 1 {
		DBStringConnect = dbStringConnectVar
	} else {
		DBStringConnect = dbStringConnect
	}

	if len(serverAddress) < 2 {
		serverAddress = serverAddressVar
	}
	if len(baseURL) < 9 {
		baseURL = baseURLVar
	}
	if len(fileStoragePath) < 1 {
		fileStoragePath = fileStoragePathVar
	}

	LoadConfig(serverAddress, baseURL)
	storage.ResoreDB(fileStoragePath)

}
