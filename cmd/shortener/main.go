package main

import (
	"io"
	"net/http"
	"strconv"
)

const server = "localhost"
const port = "8080"

var bd = map[int]string{}
var index int

func main() {

	http.HandleFunc("/", mainFunc)
	http.ListenAndServe(":"+port, nil)
}

func mainFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		id_input, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(bd[id_input]) > 2 {
			newUrl := bd[id_input]
			http.Redirect(w, r, newUrl, http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		return
	} else if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		index = len(bd)
		bd[index] = string(b)

		w.Header().Set("content-type", "http")
		w.WriteHeader(http.StatusCreated)

		w.Write([]byte(server + ":" + port + "/?id=" + strconv.Itoa(index)))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
