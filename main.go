package main

import (
	"os"
	"fmt"
	"github.com/vvotm/webimgspider/app"
)

func main()  {
	args := os.Args
	argsLength := len(args)
	if argsLength == 1 {
		fmt.Println(`
	Use: webimgspider {url} [savepath] [totalpage]
	url: website which you want fetch
	savepath: the path image will save
	totalpage: it will replace {page} to fetch image.

		ie: http://www.zhuangbi.info?page={page}
		the {} page will replace by [totalpage]
		`)
		os.Exit(1)
	}

	url := args[1]
	savepath := "./tmp"
	totalpage := "0"
	if argsLength >= 3 {
		savepath = args[2]
	}
	if argsLength >= 4 {
		totalpage = args[3]
	}
	app.Run(url, savepath, totalpage)
}
