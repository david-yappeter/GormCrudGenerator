package dir

import (
	"fmt"
	"os"
)

func Mkdir(path string) {
	if err := os.Mkdir(path, os.ModeDir); err != nil && !os.IsExist(err) && !os.IsNotExist(err) {
		panic(err)
	}
}

func FileCreate(path string, name string) *os.File {
	file, err := os.Create(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(path), name))
	if err != nil {
		panic(err)
	}
	return file
}

func pathRemoveLastStrip(path string) string {
	if len(path) > 0 && path[len(path)-1] == []byte("/")[0] {
		return path[0 : len(path)-1]
	}
	return path
}
