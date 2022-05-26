package handlers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func SetCookies(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//		fmt.Println(`GET`)

		_, err := r.Cookie("login")
		if err != nil {

			//	fmt.Println(`ПУсто`)
			token := getToken(24)
			//			fmt.Println(`SET`)
			cookie := http.Cookie{
				Name:  "login",
				Value: token,
				Path:  "/",
			}

			http.SetCookie(w, &cookie)
			r.AddCookie(&cookie)
		} else {
			//	fmt.Println(`НЕПусто` + log1.Value)
			/*
				status := storage.CheckLoginDB(log1.Value)
				fmt.Println(status)
				if status != "Y" {

					token := getToken(24)
					//			fmt.Println(`SET`)
					cookie := http.Cookie{
						Name:  "login",
						Value: token,
						Path:  "/",
					}

					http.SetCookie(w, &cookie)
					r.AddCookie(&cookie)

				}*/
		}

		h.ServeHTTP(w, r)

	})
}

func Ungzip(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		urlP, err := url.Parse(string(b))
		if err != nil {
			zr, _ := gzip.NewReader(bytes.NewReader(b))
			var b2 bytes.Buffer
			b2.ReadFrom(zr)
			zr.Close()

			urlP, err = url.Parse(b2.String())
			if err != nil {
				fmt.Println(`err - parsing url b2`)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(urlP.String())))
		h.ServeHTTP(w, r)

	})
}

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

		login, err := r.Cookie("login")
		if err != nil {
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

		intOut, err := storage.PutDB(login.Value, urlP.String())
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

func GetAPIJSONLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	login, err := r.Cookie("login")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	map1 := make(map[int]string)
	map1 = storage.GetDBLogin(login.Value)

	//	fmt.Println(map1)

	if len(map1) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	type OutData struct {
		Short_url    string `json:"short_url"`
		Original_url string `json:"original_url"`
	}
	type linksData []OutData
	var links linksData

	//links := []OutData{}
	for k := range map1 {

		links = append(links, OutData{Short_url: strconv.Itoa(k), Original_url: map1[k]})
	}

	j, err := json.Marshal(links)

	if err != nil {
		fmt.Println(`err-marshal`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	}

	return
}

func MainHandlFuncPost(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len((b)) < 10 {
		fmt.Println(`URL -no correct`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		urlP, err := url.Parse(string(b))
		if err != nil {
			zr, _ := gzip.NewReader(bytes.NewReader(b))
			var b2 bytes.Buffer
			b2.ReadFrom(zr)
			zr.Close()

			urlP, err = url.Parse(b2.String())
			if err != nil {
				fmt.Println(`err - parsing url b2`)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	*/

	login, err := r.Cookie("login")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	urlP, err := url.Parse(string(b))
	if err != nil {
		fmt.Println(`err - parsing url b2`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intOut, err := storage.PutDB(login.Value, urlP.String())
	if err != nil {
		fmt.Println(`err storage storage.DataPut`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "http")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(MakeString(strconv.Itoa(intOut))))
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
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
