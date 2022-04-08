package model_test

import (
	"dcardHw/src/model"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	e := model.Init()
	if e == nil {
		t.Log("Init success")
	} else {
		t.Error("Init fail")
	}
}

func TestGetCounter(t *testing.T) {
	count := model.GetCounter()
	if count >= 1 {
		t.Log("Get counter susses")
	} else {
		t.Error("Get counter fail")
	}
}

func TestUpdateCounter(t *testing.T) {
	count := model.GetCounter()
	if count >= 1 {
		t.Log("Get counter susses")
	} else {
		t.Error("Get counter fail")
	}
	model.UpdateCounter()
	if model.GetCounter()-count == 1 {
		t.Log("Update counter success")
	} else {
		t.Error("Update counter fail")
	}
}
func TestInsertAndGet(t *testing.T) {
	testTime := time.Now().Add(time.Hour)
	dt := time.Duration(testTime.Sub(time.Now()))
	model.SetShortenUrl("testOri", ("short" + testTime.Format("2006-01-02 15:04:05")), testTime.Format("2006-01-02 15:04:05"), dt)
	ori := model.GetOriUrl(("short" + testTime.Format("2006-01-02 15:04:05")))
	if ori == "testOri" {
		t.Log("Insert and get success")
	} else {
		t.Error("Insert and get fail")
	}
}

func TestGetShort(t *testing.T) {
	testTime := time.Now().Add(time.Hour)
	dt := time.Duration(testTime.Sub(time.Now()))
	model.SetShortenUrl("testOri", "short1"+testTime.Format("2006-01-02 15:04:05"), testTime.Format("2006-01-02 15:04:05"), dt)
	short := model.GetShortbyOri("["+testTime.Format("2006-01-02 15:04:05"), "["+testTime.Format("2006-01-02 15:04:05"))[0]
	if short == testTime.Format("2006-01-02 15:04:05") {
		t.Log("Get Short success")
	} else {
		t.Error("Get Short fail")
	}
}
