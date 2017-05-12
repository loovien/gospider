package app

import (
	"regexp"
	"strings"
	"strconv"
	"github.com/vvotm/webimgspider/except"
	"sync"
	"os"
	"math/rand"
	"github.com/vvotm/webimgspider/app/typeMatch"
	"github.com/vvotm/webimgspider/app/utils"
	"fmt"
)


var wg sync.WaitGroup

func Run(url string, path string, page string) {
	regex := regexp.MustCompile("\\{page\\}")
	newPage, err := strconv.Atoi(page)
	except.ErrorHandler(err)
	if newPage > 0 && regex.FindAllString(url, -1) != nil {
		for i := 1; i <= newPage ; i++  {
			wg.Add(1)
			pageIndex := strconv.Itoa(i)
			fmt.Println(pageIndex)
			newUrl := strings.Replace(url, "{page}", pageIndex, -1)
			fmt.Println(newUrl)
			go fetch(url, path)
		}
	} else {
		wg.Add(1)
		go fetch(url, path)
	}
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

	content := utils.FetchUrl(url)
	types := []string{"jpg", "png", "gif"}
	resultSet := typeMatch.GetImgSlice(content, -1, types)

	for _, url := range(resultSet)  {
		imgStr := utils.FetchUrl(url)
		lastSlashIndex := strings.LastIndex(url, "/")
		filename := strings.Trim(path, "/") + "/" + url[lastSlashIndex+1:]
		utils.WriteFile(imgStr, filename)
	}
	wg.Done()
}