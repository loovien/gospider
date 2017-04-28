package app

import (
	"regexp"
	"strings"
	"strconv"
	"github.com/vvotm/webimgspider/except"
	"sync"
	"os"
	"math/rand"
	"net/http"
	"io/ioutil"
)

var wg sync.WaitGroup

func Run(url string, path string, page string) {
	regex := regexp.MustCompile("\\{page\\}")

	page, err = strconv.Atoi(page)
	except.ErrorHandler(err)

	if page > 0 && regex.FindAllString(url, -1) != nil {
		for i := 1; i <= page ; i++  {
			wg.Add(1)
			pageIndex := strconv.Itoa(i)
			url = strings.Replace(url, "\\{page\\}", pageIndex , 0)
			go fetch(url, path)
		}
	}

	go fetch(url, path)
	wg.Wait()
}

func fetch(url, path string)  {

	imgname := strconv.Itoa(rand.Int())

	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex != -1 {
		imgname = url[lastSlashIndex+1:]
	}

	lastQutesIndex := strings.LastIndex(imgname, "?")
	if lastQutesIndex != -1 {
		imgname = imgname[:lastQutesIndex]
	}
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(path, 755)
	}

	resp, err := http.Get(url)
	except.ErrorHandler(err)
	defer resp.Body.Close()

	file, err := os.Create(imgname)
	defer file.Close()
	except.ErrorHandler(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	except.ErrorHandler(err)
	file.Write(bytes)

	wg.Done()
}