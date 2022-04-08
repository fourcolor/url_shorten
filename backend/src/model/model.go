package model

import (
	redis "dcardHw/src/model/redis"
	"time"
)

func SetShortenUrl(ori, shorten, zv string, t time.Duration) {
	redis.SetShortenUrl(ori, shorten, zv, t)
}

func GetShortbyOri(min, max string) (url []string) {
	url = redis.GetShortbyOri(min, max)
	return
}

func GetOriUrl(short string) string {
	return redis.GetOriUrl(short)
}

func GetCounter() (val int64) {
	return redis.GetCounter()
}

func UpdateCounter() {
	redis.UpdateCounter()
}

func Init() (e error) {
	e = redis.Init()
	return
}
