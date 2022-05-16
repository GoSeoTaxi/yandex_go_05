package handlers

import (
	"bytes"
	"compress/flate"
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

func MainHandlFuncPost(w http.ResponseWriter, r *http.Request) {

	//oplog := httplog.LogEntry(r.Context())
	//oplog.Printf(http.MethodPost)
	//	oplog.Printf(r.Body)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//	if len(string(b)) < 10 {
	//		w.WriteHeader(http.StatusBadRequest)
	//		return
	//	}

	fmt.Println(`++++++++++++++++++++++`)
	fmt.Println()
	fmt.Println(`++++++++++++++++++++++`)

	urlP, err := url.Parse(string(b))
	if err != nil {
		fmt.Println(`++++++++++++`)
		fmt.Println(`err - но parsing`)
		fmt.Println(b)
		fmt.Println(string(b))
		fmt.Println(`++++++++++++`)
		rea := flate.NewReader(bytes.NewReader(b))
		defer rea.Close()
		var b2 bytes.Buffer
		// в переменную b записываются распакованные данные
		_, err := b2.ReadFrom(rea)
		if err != nil {
			fmt.Println(err)
			fmt.Println(`err - decompress`)
		}
		fmt.Println(b2.Bytes())
		fmt.Println(b2.String())
		urlP, err = url.Parse(b2.String())
		if err != nil {
			fmt.Println(`err - parsing url b2`)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	//		a := storage.DataPut{URL1: urlP.String()}

	intOut, err := storage.PutDB(urlP.String())
	if err != nil {
		fmt.Println(`err storage storage.DataPut`)
	}

	w.Header().Set("content-type", "http")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(MakeString(strconv.Itoa(intOut))))
	return

}

func MainHandlFuncGet(w http.ResponseWriter, r *http.Request) {

	//oplog := httplog.LogEntry(r.Context())
	//oplog.Printf(http.MethodPost)
	//	oplog.Printf(r.Body)

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

}
