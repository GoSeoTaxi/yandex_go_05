package handlers

import (
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/storage"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func MainHandlFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idInput, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idGetQuery := storage.DataGet{IDURLRedirect: idInput}
		urlOut2redir, err := idGetQuery.GetDB()
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

		if len(string(b)) < 10 &&
			!strings.Contains(string(b), ".") &&
			!strings.Contains(string(b), "://") &&
			!strings.Contains(string(b), "http") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		a := storage.DataPut{URL1: string(b)}
		intOut, err := a.PutDB()
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
