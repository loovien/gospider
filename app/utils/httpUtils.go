package utils

import (
	"net/http"
	"github.com/vvotm/webimgspider/except"
	"io/ioutil"
)

func FetchUrl(url string) (content string){
	if url[:4] != "http" {
		url = "http:" + url
	}
	resp, err := http.Get(url)
	except.ErrorHandler(err)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	except.ErrorHandler(err)

	return string(bytes)

}
