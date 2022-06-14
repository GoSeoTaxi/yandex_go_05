package config

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	_ "github.com/lib/pq"
	"os"
)

func InitCLI() {

	var serverAddress string
	var baseURL string
	var fileStoragePath string
	var dbStringConnect string

	flag.StringVar(&dbStringConnect, "d", "user=postgres password=qwerty dbname=goyp sslmode=disable", "DATABASE_DSN")
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

	//check DB
	db, err := sql.Open("postgres", DBStringConnect)
	if err != nil {
		fmt.Println(`StaticTest - err`)
	}
	defer db.Close()
	ping := db.Ping()

	if len(DBStringConnect) > 1 && err == nil && ping == nil {
		_, err := db.Exec("create table if not exists shortyp10 (id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,link  varchar(256) NOT NULL,login varchar(256),is_del boolean default false);")
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("create unique index if not exists shortyp10_link_uindex\n    on shortyp10 (link);")
		if err != nil {
			fmt.Println(err)
		}

		_, err = db.Exec("INSERT INTO shortyp10 (link, login, is_del) 	VALUES ('0', '0', DEFAULT);")
		//	"create unique index if not exists shortyp10_link_uindex\n    on shortyp10 (link);")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(`use db`)
		etc.UseDB = "Y"

	} else {
		fmt.Println(`DB - DIE`)
		storage.ResoreDB(fileStoragePath)
	}

	storage.StringConnect = DBStringConnect
	LoadConfig(serverAddress, baseURL)

}
