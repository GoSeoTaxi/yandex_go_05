package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func APIJSON(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		type urlInputJSON struct {
			URL string `json:"url"`
		}
		var apiJSONInput urlInputJSON
		err = json.Unmarshal(b, &apiJSONInput)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		urlP, err := url.Parse(apiJSONInput.URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(urlP.String()) < 10 ||
			!json.Valid([]byte(b)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//		if len(apiJsonInput.Url) < 10 ||
		//			!json.Valid([]byte(b)) ||
		//			!strings.Contains(apiJsonInput.Url, ".") ||
		//			!strings.Contains(apiJsonInput.Url, "://") ||
		//			!strings.Contains(apiJsonInput.Url, "http") {
		//			w.WriteHeader(http.StatusBadRequest)
		//			return
		//		}

		//	a := storage.DataPut{URL1: urlP.String()}
		intOut, err := storage.PutDB(urlP.String())
		//	intOut, err := a.PutDB()
		if err != nil {
			fmt.Println(`err storage storage.DataPut`)
		}

		urlOut := MakeString(strconv.Itoa(intOut))
		urlOutMap := map[string]string{
			"result": urlOut,
		}
		urlOutByte, err := json.Marshal(urlOutMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(urlOutByte)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func MainHandlFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idInput, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Println(`ERR - GET id`)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//	idGetQuery := storage.DataGet{IDURLRedirect: idInput}
		//	urlOut2redir, err := idGetQuery.GetDB()
		urlOut2redir, err := storage.GetDB(idInput)
		if err != nil {
			fmt.Println(`ERR storage DataGet`)
		}

		if len(urlOut2redir) > 2 {
			http.Redirect(w, r, urlOut2redir, http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	} else if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if len(string(b)) < 10 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(`++++++`)
		fmt.Println(b)
		fmt.Println(string(b))

		urlP, err := url.Parse(string(b))
		if err != nil {
			fmt.Println(`err - parsing url`)
			//		w.WriteHeader(http.StatusBadRequest)
			return
		}

		str1 := "0"

		str1 = urlP.String()
		str1 = string(b)

		//		a := storage.DataPut{URL1: urlP.String()}

		intOut, err := storage.PutDB(str1)
		if err != nil {
			fmt.Println(`err storage storage.DataPut`)
		}

		w.Header().Set("content-type", "http")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(MakeString(strconv.Itoa(intOut))))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
