package routes_test

import (
	"bytes"
	api "dcardHw/src/api"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ShortenerReq struct {
	Url      string    `json:"url" form:"url" `
	ExpireAt time.Time `json:"expireAt" form:"expireAt" `
}

func TestShortenerforPositive(t *testing.T) {
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
	fmt.Println(w)
	if w.Code == http.StatusAccepted {
		t.Log("ShortenerforPositive success")
	} else {
		t.Error("ShortenerforPositive fail")
	}
}

func TestShortenerforNegative(t *testing.T) {
	server := api.StartServer()
	data := map[string]string{
		"uUrl":     "https://www.google.com",
		"ExpireAt": "aaaaa",
	}
	jsonValue, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/v1/urls", bytes.NewBuffer(jsonValue))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code == http.StatusBadRequest {
		t.Log("ShortenerforNegative validator success")
	} else {
		t.Error("ShortenerforNegative validator field fail")
	}
	data = map[string]string{
		"url":      "aaaa",
		"expireAt": "2023-02-08T09:20:41Z",
	}
	jsonValue, _ = json.Marshal(data)
	req, err = http.NewRequest("POST", "http://127.0.0.1:8080/api/v1/urls", bytes.NewBuffer(jsonValue))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code == http.StatusBadRequest {
		t.Log("ShortenerforNegative invalid url success")
	} else {
		t.Error("ShortenerforNegative invalid url fail")
	}
}
