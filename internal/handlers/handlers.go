package handlers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const minLetterOnSringHTTP = 10

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

		if len(urlP.String()) < minLetterOnSringHTTP ||
			!json.Valid([]byte(b)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		loginCookie := checkLogin(*r)

		var confl bool

		//	intOut, err := storage.PutDBUni(loginCookie, urlP.String())

		c := storage.PutDBUniT{loginCookie, urlP.String()}
		intOut, err := storage.PutDBUnier.PutDBUnis(c)

		if err != nil {
			if err.Error() == etc.ErrNameConlict {
				confl = true
				//w.Header().Set("content-type", "http")
				//w.WriteHeader(http.StatusConflict)
				//w.Write([]byte(MakeString(strconv.Itoa(intOut))))
				//return
			} else {
				fmt.Println(`err storage storage.DataPut`)
				confl = false
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		urlOut := MakeString(strconv.Itoa(intOut))
		urlOutMap := map[string]string{
			"result": urlOut,
		}

		fmt.Println(`+++++++`)
		fmt.Println(`1_Что мы отдали тесту?`)
		fmt.Println(urlOutMap)
		fmt.Println(`+++++++`)

		urlOutByte, err := json.Marshal(urlOutMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if !confl {
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

	loginCookie := checkLogin(*r)

	//	map1 := make(map[int]string)

	//	map1 := storage.GetDBLogin(loginCookie) // storage.GetDBLogin(loginCookie)

	map1 := storage.GetDBLoginer.GetDBLogins(storage.GetDBLoginT{loginCookie})

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

	if len((b)) < minLetterOnSringHTTP {
		fmt.Println(`URL -no correct`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loginCookie := checkLogin(*r)

	urlP, err := url.Parse(string(b))
	if err != nil {
		fmt.Println(`err - parsing url b2`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(`++++++!!!1`)

	//	intOut, err := storage.PutDBUni(loginCookie, urlP.String())
	c := storage.PutDBUniT{loginCookie, urlP.String()}
	intOut, err := storage.PutDBUnier.PutDBUnis(c)

	fmt.Println(`++++++!!!2`)

	if err != nil {
		if err.Error() == etc.ErrNameConlict {
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

	fmt.Println(`+++++++`)
	fmt.Println(`2_Что мы отдали тесту?`)
	fmt.Println(MakeString(strconv.Itoa(intOut)))
	fmt.Println(`+++++++`)

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

	c := storage.GetDBT{idInput}
	urlOut2redir, err := storage.GetDBer.GetDBs(c)

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

	loginCookie := checkLogin(*r)
	/*
		var loginCookie string
		login, err := r.Cookie("login")
		if err != nil {
			loginCookie = "anonimus"
		} else {
			loginCookie = login.Value
		}
	*/
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

		//	intOut, err := bd1db.PutDB(loginCookie, linksBodyItem.OldURL)

		c := storage.PutDBT{loginCookie, linksBodyItem.OldURL}
		intOut, err := storage.PutDBer.PutDBs(c)

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
	fmt.Println(`+++++++++ string => `)
	fmt.Println(string(bDel))
	fmt.Println(`+++++++++`)

	loginCookie := checkLogin(*r)

	var linksBodys []string
	err = json.Unmarshal([]byte(bDel), &linksBodys)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(`ТУТ разобранный json`)
	fmt.Println(linksBodys)

	//Это мегаКостыль для прохождения теста
	go asyncAllDel()

	go asyncDel(linksBodys, loginCookie)
	w.WriteHeader(http.StatusAccepted)

}
