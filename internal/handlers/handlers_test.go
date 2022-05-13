package handlers

import (
	"bytes"
	"encoding/json"
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
	stringRequest1 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	MainHandlFunc(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t1.TestMainHandlFunc`)
	}

	//проверка post Запроса

	stringRequest2 := config.ServerHost + ":" + config.Port + "/"

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
		fmt.Println(`err t2.TestMainHandlFunc`)
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
		fmt.Println(`err t3.TestMainHandlFunc`)
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
		fmt.Println(`err t4.TestMainHandlFunc`)
	}

	// проверка результатов подтеста
	if cointTests != tests {
		t.Fatalf("testing is not ok coint=%v , need=%v", cointTests, tests)
	} else {
		fmt.Println(`TEST-OK-HANDLERS-MAIN`)
	}

}

func TestApiJson(t *testing.T) {

	const testsTestApiJson = 4
	urlTestTestApiJson := `{"url":"https://gmail.com"}`
	urlTestTestApiJsonReq := "https://gmail.com"
	urlTestTestApiRequest := `1` //переменная для теста json

	cointTestsTestApiJson := 0

	//проверка на ошибку по пустой базе
	idTest := "1"
	stringRequest1 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	ApiJson(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTestsTestApiJson += 1
	} else {
		fmt.Println(`err t1.TestApiJson`)
	}

	//проверка post Запроса
	stringRequest2 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest

	urlTestTestApiRequest = `{"result":"` + stringRequest2 + `"}`

	buffer := new(bytes.Buffer)

	buffer.WriteString(urlTestTestApiJson)

	req2, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req2.Header.Set("content-type", "application/json")
	if err != nil {
		fmt.Println(`err t2.TestApiJson!`)
		t.Fatalf("not req :%v", err)
	}
	rec2 := httptest.NewRecorder()
	ApiJson(rec2, req2)
	t21 := rec2.Body

	if t21.String() == urlTestTestApiRequest && rec2.Code == http.StatusCreated {
		cointTestsTestApiJson += 1
	} else {
		fmt.Println(`err t2.TestApiJson!!`)
	}

	//проверка get Запроса по результату post ответа запроса

	stringRequest3Byte := t21.Bytes()

	type resJSON struct {
		Url string `json:"result"`
	}
	var apiJsonInput resJSON
	err = json.Unmarshal(stringRequest3Byte, &apiJsonInput)
	if err != nil {
		fmt.Println(`err t3.TestApiJson!`)
		return
	}

	stringRequest3 := apiJsonInput.Url
	req3, err := http.NewRequest(http.MethodGet, stringRequest3, nil)
	if err != nil {
		fmt.Println(`err t3.TestApiJson!!`)
		t.Fatalf("not req :%v", err)
	}
	rec3 := httptest.NewRecorder()
	MainHandlFunc(rec3, req3)

	//	strEq1 := "<a href=" + url.QueryEscape(urlTestTestApiJsonReq) + "\">Temporary Redirect</a>."
	//Вот тут нужно почитать по подробнее
	strEq1 := "<a href=\"" + urlTestTestApiJsonReq + "\">Temporary Redirect</a>."
	t31 := rec3.Body

	if rec3.Code == http.StatusTemporaryRedirect && strings.Contains(t31.String(), strEq1) == true {
		cointTestsTestApiJson += 1
	} else {
		fmt.Println(`err t3.TestApiJson!!!`)
	}

	//проверка на пустой post запрос

	buffer4 := new(bytes.Buffer)
	params4 := url.Values{}

	buffer4.WriteString(params4.Encode())

	req4, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req4.Header.Set("content-type", "application/json")
	if err != nil {
		fmt.Println(`err t4.TestApiJson!!!`)
		t.Fatalf("not req :%v", err)
	}
	rec4 := httptest.NewRecorder()
	ApiJson(rec4, req4)

	if rec4.Code == http.StatusBadRequest {
		cointTestsTestApiJson += 1
	} else {
		fmt.Println(`err t4.TestApiJson`)
	}

	// проверка результатов подтеста
	if cointTestsTestApiJson != testsTestApiJson {
		t.Fatalf("testing is not ok coint=%v , need=%v", cointTestsTestApiJson, testsTestApiJson)
	} else {
		fmt.Println(`TEST-OK-HANDLERS-JSON`)
	}

}
