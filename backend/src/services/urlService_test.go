package services_test

import (
	model "dcardHw/src/model"
	"dcardHw/src/services"
	"testing"
	"time"
)

var testOriUrl = "https://www.google.com"

var testTime = time.Now().Add(time.Hour)

func TestGenerateShortenUrlPositive(t *testing.T) {
	model.Init()
	status, shorturl := services.GenerateShortenUrl(testOriUrl, testTime)
	first := shorturl
	if status == 0 {
		t.Log("Insert first success")
	} else {
		t.Error("Insert first fail")
	}
	status, shorturl = services.GenerateShortenUrl(testOriUrl, testTime)
	if status == 0 && shorturl == first {
		t.Log("Insert again success")
	} else {
		t.Error("Insert again fail")
	}
}

func TestGenerateShortenUrlNegative(t *testing.T) {
	model.Init()
	badUrl := "AAaaa"
	status, shorturl := services.GenerateShortenUrl(badUrl, testTime)
	if status == 1 && shorturl == "" {
		t.Log("Insert bad ori url success")
	} else {
		t.Error("Insert bad ori url fail")
	}
}

func TestRedirectUrlPositive(t *testing.T) {
	model.Init()
	status, shorturl := services.GenerateShortenUrl(testOriUrl, testTime)
	if status == 0 {
		model.UpdateCounter()
		t.Log("Insert first success")
	} else {
		t.Error("Insert first fail")
	}
	status1, ori := services.RedirectUrl(shorturl)
	if status1 == 0 && ori == testOriUrl {
		t.Log("Redirect success")
	} else {
		t.Error("Redirect fail")
	}
}

func TestRedirectUrlNegative(t *testing.T) {
	model.Init()
	status, shorturl := services.GenerateShortenUrl(testOriUrl, testTime)
	if status == 0 {
		model.UpdateCounter()
		t.Log("Insert first success")
	} else {
		t.Error("Insert first fail")
	}
	status1, ori := services.RedirectUrl(shorturl + "1")
	if status1 == 0 && ori == testOriUrl {
		t.Error("Redirect fail")
	} else {
		t.Log("Redirect success")
	}
}
