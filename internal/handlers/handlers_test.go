package handlers

import (
	"bytes"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/config"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestMainHandlFunc(t *testing.T) {

	const tests = 4
	const urlTest = "https://ya.ru/"
	cointTests := 0

	//проверка на ошибку по пустой базе
	idTest := "0"
	stringRequest1 := config.Server_host + ":" + config.Port + "/?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	MainHandlFunc(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t1`)
	}

	//проверка post Запроса

	stringRequest2 := config.Server_host + ":" + config.Port + "/"

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("url", urlTest)
	buffer.WriteString(params.Encode())

	req2, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req2.Header.Set("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec2 := httptest.NewRecorder()
	MainHandlFunc(rec2, req2)

	t21 := rec2.Body
	if t21.String() == stringRequest1 && rec2.Code == http.StatusCreated {
		cointTests += 1
	} else {
		fmt.Println(`err t2`)
	}

	//проверка get Запроса по результату post ответа запроса

	stringRequest3 := t21.String()
	req3, err := http.NewRequest(http.MethodGet, stringRequest3, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec3 := httptest.NewRecorder()
	MainHandlFunc(rec3, req3)

	strEq1 := "<a href=\"/url=" + url.QueryEscape(urlTest) + "\">Temporary Redirect</a>."
	t31 := rec3.Body

	if rec3.Code == http.StatusTemporaryRedirect && strings.Contains(t31.String(), strEq1) == true {
		cointTests += 1
	} else {
		fmt.Println(`err t3`)
	}

	//проверка на пустой post запрос

	buffer4 := new(bytes.Buffer)
	params4 := url.Values{}

	buffer4.WriteString(params4.Encode())

	req4, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req4.Header.Set("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec4 := httptest.NewRecorder()
	MainHandlFunc(rec4, req4)

	if rec4.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t4`)
	}

	// проверка результатов подтеста
	if cointTests != tests {
		t.Fatalf("testing is not ok coint=%v , need=%v", cointTests, tests)
	} else {
		fmt.Println(`TEST-OK-HANDLERS`)
	}

}
