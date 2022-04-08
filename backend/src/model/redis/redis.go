package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func Init() error {
	client := Connect()
	defer client.Close()
	tv, err := client.Get("total").Result()
	if err == redis.Nil {
		err = client.Set("total", 1, 0).Err()
		if err != nil {
			return err
		}
	}
	var val int64
	_, err = fmt.Sscan(tv, &val)
	return err
}

func GetShortbyOri(min, max string) []string {
	client := Connect()
	defer client.Close()
	r := redis.ZRangeBy{Min: min, Max: max, Offset: 0, Count: 1}
	data := client.ZRangeByLex("url", r)
	fmt.Println(data)
	res, _ := data.Result()
	return res
}

func GetOriUrl(shortenUrl string) string {
	client := Connect()
	defer client.Close()
	val, err := client.Get(shortenUrl).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	if err == redis.Nil {
		return ""
	}
	return val
}

func SetShortenUrl(ori, shorten, zv string, t time.Duration) {
	client := Connect()
	defer client.Close()
	err := client.Set(shorten, ori, t).Err()
	if err != nil {
		panic(err)
	}
	z := redis.Z{Score: 0, Member: zv}
	client.ZAdd("url", z)
}

func GetCounter() (val int64) {
	client := Connect()
	defer client.Close()
	tv, _ := client.Get("total").Result()
	fmt.Sscan(tv, &val)
	return
}

func UpdateCounter() {
	client := Connect()
	client.Incr("total").Val()
}
