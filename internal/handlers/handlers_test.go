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

func TestAPIJSON(t *testing.T) {

	config.LoadConfig("", "")

	const testsTestAPIJSON = 4
	urlTestTestAPIJSON := `{"url":"https://gmail.com"}`
	urlTestTestAPIJSONReq := "https://gmail.com"
	urlTestTestAPIRequest := `1` //переменная для теста json

	cointTestsTestAPIJSON := 0

	//проверка на ошибку по пустой базе
	idTest := "0"
	stringRequest1 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	APIJSON(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTestsTestAPIJSON += 1
	} else {
		fmt.Println(`err t1.TestApiJson`)
	}

	//проверка post Запроса
	stringRequest2 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest

	urlTestTestAPIRequest = `{"result":"` + stringRequest2 + `"}`

	buffer := new(bytes.Buffer)

	buffer.WriteString(urlTestTestAPIJSON)

	req2, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req2.Header.Set("content-type", "application/json")
	if err != nil {
		fmt.Println(`err t2.TestApiJson!`)
		t.Fatalf("not req :%v", err)
	}
	rec2 := httptest.NewRecorder()
	APIJSON(rec2, req2)
	t21 := rec2.Body

	if t21.String() == urlTestTestAPIRequest && rec2.Code == http.StatusCreated {
		cointTestsTestAPIJSON += 1
	} else {
		fmt.Println(`err t2.TestApiJson!!`)
	}

	//проверка get Запроса по результату post ответа запроса
	type resJSON struct {
		URL string `json:"result"`
	}
	var apiJSONInput resJSON

	errJ := json.Unmarshal(t21.Bytes(), &apiJSONInput)
	if errJ != nil {
		fmt.Println(`err t3.TestApiJson!`)
	}

	req3, err1 := http.NewRequest(http.MethodGet, apiJSONInput.URL, nil)
	if err1 != nil {
		fmt.Println(`err t3.TestApiJson!!`)
	}
	rec3 := httptest.NewRecorder()
	MainHandlFuncGet(rec3, req3)

	//	strEq1 := "<a href=" + url.QueryEscape(urlTestTestAPIJSONReq) + "\">Temporary Redirect</a>."
	//Вот тут нужно почитать по подробнее
	strEq1 := "<a href=\"" + urlTestTestAPIJSONReq + "\">Temporary Redirect</a>."
	t31 := rec3.Body

	if rec3.Code == http.StatusTemporaryRedirect && strings.Contains(t31.String(), strEq1) == true {
		cointTestsTestAPIJSON += 1
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
	APIJSON(rec4, req4)

	if rec4.Code == http.StatusBadRequest {
		cointTestsTestAPIJSON += 1
	} else {
		fmt.Println(`err t4.TestApiJson`)
	}

	// проверка результатов подтеста
	if cointTestsTestAPIJSON != testsTestAPIJSON {
		t.Fatalf("testing is not ok coint=%v , need=%v", cointTestsTestAPIJSON, testsTestAPIJSON)
	} else {
		fmt.Println(`TEST-OK-HANDLERS-JSON`)
	}

}

func TestMainHandlFuncPost(t *testing.T) {

	const testsPost = 4
	const urlTestPost = "https://ya.ru/"
	cointTests := 0

	fmt.Println(`++++++++!`)
	//проверка на ошибку по пустой базе
	idTest := "1"
	stringRequest1 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	MainHandlFuncGet(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t1.TestMainHandlFunc`)
	}

	//проверка post Запроса

	fmt.Println(`++++++++!!`)

	stringRequest2 := config.ServerHost + ":" + config.Port + config.PathURLConf

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("url", urlTestPost)
	buffer.WriteString(params.Encode())

	req2, err := http.NewRequest(http.MethodPost, stringRequest2, buffer)
	req2.Header.Set("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec2 := httptest.NewRecorder()
	MainHandlFuncPost(rec2, req2)

	t21 := rec2.Body

	if t21.String() == stringRequest1 && rec2.Code == http.StatusCreated {
		cointTests += 1
	} else {
		fmt.Println(`err t2.TestMainHandlFunc`)
	}

	fmt.Println(`++++++++!!!`)

	//проверка get Запроса по результату post ответа запроса
	stringRequest3 := t21.String()
	req3, err := http.NewRequest(http.MethodGet, stringRequest3, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec3 := httptest.NewRecorder()
	MainHandlFuncGet(rec3, req3)

	strEq1 := "<a href=\"/url=" + url.QueryEscape(urlTestPost) + "\">Temporary Redirect</a>."
	t31 := rec3.Body

	if rec3.Code == http.StatusTemporaryRedirect && strings.Contains(t31.String(), strEq1) == true {
		cointTests += 1
	} else {
		fmt.Println(`err t3.TestMainHandlFunc`)
	}

	fmt.Println(`++++++++!!!!`)
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
	MainHandlFuncPost(rec4, req4)

	if rec4.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t4.TestMainHandlFunc`)
	}

	// проверка результатов подтеста
	if cointTests != testsPost {
		t.Fatalf("testing is not ok coint=%v , need=%v", cointTests, testsPost)
	} else {
		fmt.Println(`TEST-OK-HANDLERS-MAIN`)
	}

}

/*

func TestMainHandlFuncGet(t *testing.T) {

	const tests = 4
	const urlTest = "https://ya.ru/"
	cointTests := 0

	//проверка на ошибку по пустой базе
	idTest := "1"
	stringRequest1 := config.ServerHost + ":" + config.Port + config.PathURLConf + "?" + config.ConstGetEndPoint + "=" + idTest
	req1, err := http.NewRequest(http.MethodGet, stringRequest1, nil)
	if err != nil {
		t.Fatalf("not req :%v", err)
	}
	rec1 := httptest.NewRecorder()
	MainHandlFuncGet(rec1, req1)
	if rec1.Code == http.StatusBadRequest {
		cointTests += 1
	} else {
		fmt.Println(`err t1.TestMainHandlFunc`)
	}

	//проверка post Запроса

	stringRequest2 := config.ServerHost + ":" + config.Port + config.PathURLConf

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
	MainHandlFuncGet(rec2, req2)

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
	MainHandlFuncGet(rec3, req3)

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
	MainHandlFuncGet(rec4, req4)

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


*/
