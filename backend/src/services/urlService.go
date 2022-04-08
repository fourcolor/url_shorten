package services

import (
	model "dcardHw/src/model"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GenerateShortenUrl(ori string, expireAt time.Time) (status int, shortenUrl string) {
	if time.Now().After(expireAt) {
		status = 2
		return
	}
	_, err := url.ParseRequestURI(ori)
	status = 0
	if err != nil {
		status = 1
		return
	}
	t := expireAt.Unix()
	min := "[" + strconv.FormatInt(t, 10) + "#" + ori + "#"
	max := "[" + strconv.FormatInt(t, 10) + "#" + ori + "#" + "\xff"
	fmt.Println(min, max)
	res := model.GetShortbyOri(min, max)
	if len(res) == 0 {
		id := model.GetCounter()
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(id))
		shortenUrl = base64.StdEncoding.EncodeToString(b)
		dt := time.Duration(expireAt.Sub(time.Now()))
		model.SetShortenUrl(ori, shortenUrl, strconv.FormatInt(t, 10)+"#"+ori+"#"+strconv.FormatInt(id, 10), dt)
		model.UpdateCounter()
	} else {
		id, err := strconv.ParseInt(strings.Split(res[0], "#")[2], 10, 64)
		if err != nil {
			panic(err)
		}
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(id))
		shortenUrl = base64.StdEncoding.EncodeToString(b)
	}
	return

}

func RedirectUrl(shortenUrl string) (s int, ori string) {
	ori = model.GetOriUrl(shortenUrl)
	s = 0
	if ori == "" {
		s = 1
	}
	return
}
