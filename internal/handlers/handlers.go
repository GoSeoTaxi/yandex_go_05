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

		//	fmt.Println(`GET cookies`)

		_, err := r.Cookie("login")
		if err != nil {

			//	fmt.Println(`empty  cookies`)
			token := getToken(24)
			//	fmt.Println(`SET  cookies`)
			cookie := http.Cookie{
				Name:  "login",
				Value: token,
				Path:  "*",
			}

			http.SetCookie(w, &cookie)
			r.AddCookie(&cookie)
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

		var loginCookie string
		login, err := r.Cookie("login")
		if err != nil {
			//	w.WriteHeader(http.StatusBadRequest)
			//	return
			loginCookie = "anonimus"
		} else {
			loginCookie = login.Value
		}
		var confl int
		intOut, err := storage.PutDBUni(loginCookie, urlP.String())
		if err != nil {
			if err.Error() == "Conflict" {

				confl = 1
				//w.Header().Set("content-type", "http")
				//w.WriteHeader(http.StatusConflict)
				//w.Write([]byte(MakeString(strconv.Itoa(intOut))))
				//return

			} else {
				fmt.Println(`err storage storage.DataPut`)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
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
		if confl != 1 {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

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

	var loginCookie string
	login, err := r.Cookie("login")
	if err != nil {
		loginCookie = "anonimus"
	} else {
		loginCookie = login.Value
	}

	//	map1 := make(map[int]string)
	map1 := storage.GetDBLogin(loginCookie)

	if len(map1) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	type OutData struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}
	type linksData []OutData
	var links linksData

	//links := []OutData{}
	for k := range map1 {

		links = append(links, OutData{ShortURL: MakeString(strconv.Itoa(k)), OriginalURL: map1[k]})
	}

	j, err := json.Marshal(links)
	if err != nil {
		fmt.Println(`err-marshal`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	}
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

	var loginCookie string
	login, err := r.Cookie("login")
	if err != nil {
		loginCookie = "anonimus"
	} else {
		loginCookie = login.Value
	}

	urlP, err := url.Parse(string(b))
	if err != nil {
		fmt.Println(`err - parsing url b2`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intOut, err := storage.PutDBUni(loginCookie, urlP.String())

	if err != nil {
		if err.Error() == "Conflict" {
			w.Header().Set("content-type", "http")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(MakeString(strconv.Itoa(intOut))))
			return
		} else {
			fmt.Println(`err storage storage.DataPut`)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}

	w.Header().Set("content-type", "http")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(MakeString(strconv.Itoa(intOut))))
}

func MainHandlFuncGet(w http.ResponseWriter, r *http.Request) {

	idInput, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println(`ERR - GET id`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	urlOut2redir, err := storage.GetDB(idInput)
	if err != nil {
		if err.Error() == "410" {
			w.WriteHeader(http.StatusGone)
			return
		} else {
			fmt.Println(`ERR storage DataGet`)
		}
	}

	if len(urlOut2redir) > 2 {
		http.Redirect(w, r, urlOut2redir, http.StatusTemporaryRedirect)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func APIJSONBatch(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var loginCookie string
	login, err := r.Cookie("login")
	if err != nil {
		loginCookie = "anonimus"
	} else {
		loginCookie = login.Value
	}

	type urlInputJSONLine struct {
		TempID string `json:"correlation_id"`
		OldURL string `json:"original_url"`
	}
	var linksBody []urlInputJSONLine
	err = json.Unmarshal([]byte(b), &linksBody)
	if err != nil {
		fmt.Println(err)
	}

	type urlInputJSONLineNew struct {
		TempIDNew string `json:"correlation_id"`
		NewURL    string `json:"short_url"`
	}
	var linksBodyNew []urlInputJSONLineNew

	for i := 0; i < len(linksBody); i++ {
		linksBodyItem := linksBody[i]

		var a1 urlInputJSONLineNew
		a1.TempIDNew = linksBodyItem.TempID

		intOut, err := storage.PutDB(loginCookie, linksBodyItem.OldURL)
		if err != nil {
			fmt.Println(`err storage storage.DataPut Api Batch`)
		}
		a1.NewURL = MakeString(strconv.Itoa(intOut))
		linksBodyNew = append(linksBodyNew, a1)
	}

	urlOutByte, err := json.Marshal(linksBodyNew)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(urlOutByte)

}

func APIDelBatch(w http.ResponseWriter, r *http.Request) {

	bDel, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(`++++++ JSON InBody`)
	fmt.Println(bDel)
	fmt.Println(`+++++++++`)

	var loginCookie string
	login, err := r.Cookie("login")
	if err != nil {
		loginCookie = "anonimus"
	} else {
		loginCookie = login.Value
	}

	var linksBodys []string
	err = json.Unmarshal([]byte(bDel), &linksBodys)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(`ТУТ разобранный json`)
	fmt.Println(linksBodys)

	//sendFunc
	go asyncDel(linksBodys, loginCookie)
	w.WriteHeader(http.StatusAccepted)

	/*
		type urlInputJSONLineNew struct {
			TempIDNew string `json:"correlation_id"`
			NewURL    string `json:"short_url"`
		}
		var linksBodyNew []urlInputJSONLineNew

		for i := 0; i < len(linksBody); i++ {
			linksBodyItem := linksBody[i]

			var a1 urlInputJSONLineNew
			a1.TempIDNew = linksBodyItem.TempID

			intOut, err := storage.PutDB(loginCookie, linksBodyItem.OldURL)
			if err != nil {
				fmt.Println(`err storage storage.DataPut Api Batch`)
			}
			a1.NewURL = MakeString(strconv.Itoa(intOut))
			linksBodyNew = append(linksBodyNew, a1)
		}

		urlOutByte, err := json.Marshal(linksBodyNew)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(urlOutByte)
	*/

}
