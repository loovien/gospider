package utils

import (
	"net/http"
	"github.com/vvotm/gospider/except"
	"io/ioutil"
	"fmt"
)

func FetchUrl(url string) (content string){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(url + " not ok")
		}
	}()
	if url[:4] != "http" {
		url = "http:" + url
	}
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	except.ErrorHandler(err)

	return string(bytes)

}
