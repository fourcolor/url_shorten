package routes_test

import (
	"bytes"
	api "dcardHw/src/api"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ShortenerRes struct {
	Id       string `json:"id" form:"id" `
	ShortUrl string `json:"shortUrl" form:"shortUrl" `
}

func TestPRedirectPositive(t *testing.T) {
	server := api.StartServer()

	data := map[string]string{
		"url":      "https://www.google.com",
		"expireAt": "2023-02-08T09:20:41Z",
	}
	jsonValue, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/v1/urls", bytes.NewBuffer(jsonValue))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code == http.StatusAccepted {
		t.Log("ShortenerforPositive success")
	} else {
		t.Error("ShortenerforPositive fail")
	}
	fmt.Println(w.Body.String())
	response := ShortenerRes{}

	json.Unmarshal([]byte(w.Body.String()), &response)
	fmt.Println(response.ShortUrl)
	short := response.ShortUrl
	req, err = http.NewRequest("GET", "http://127.0.0.1:8080/"+short, nil)
	if err != nil {
		t.Fatal(err)
	}
	server.ServeHTTP(w, req)
	if w.Code == http.StatusAccepted {
		t.Log("TestPRedirectPositive success")
	} else {
		t.Error("TestPRedirectPositive fail")
	}
}

func TestPRedirectNegative(t *testing.T) {
	server := api.StartServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/WWWWWWWWWWWWWWWWWWWWWWW", nil)
	if err != nil {
		t.Fatal(err)
	}
	server.ServeHTTP(w, req)
	if w.Code == http.StatusNotFound {
		t.Log("TestPRedirectPositive success")
	} else {
		t.Error("TestPRedirectPositive fail")
	}
}
