package utils

import (
	"os"
	"github.com/vvotm/gospider/except"
	"io"
)

func WriteFile(content, filename string) (ok int, err error){

	handle, err := os.Create(filename)
	except.ErrorHandler(err)
	return io.WriteString(handle, content)
}
