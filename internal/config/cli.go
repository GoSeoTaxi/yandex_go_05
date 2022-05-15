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

	flag.StringVar(&serverAddress, "a", "localhost:8080", "SERVER_ADDRESS")
	flag.StringVar(&baseURL, "b", ":8080", "BASE_URL")
	flag.StringVar(&fileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	var serverAddressVar string
	var baseURLVar string
	var fileStoragePathVar string

	serverAddressVar = os.Getenv("SERVER_ADDRESS")
	baseURLVar = os.Getenv("BASE_URL")
	fileStoragePathVar = os.Getenv("FILE_STORAGE_PATH")

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
