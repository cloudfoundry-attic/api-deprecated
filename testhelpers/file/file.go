package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	writeFlag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	writePerm = 0666
	dirPerm   = 0777
)

func TmpDir() (dir string) {
	var err error
	dir, err = ioutil.TempDir("", "")
	panicErr(err, "could not create temp dir")
	return
}

func TmpFile() (path string) {
	f, err := ioutil.TempFile("", "")
	panicErr(err, "could not create temp file")
	defer f.Close()
	return f.Name()
}

func Cwd() string {
	cwd, err := os.Getwd()
	panicErr(err, "could not get cwd")
	return cwd
}

func Copy(srcPath, destPath string) {
	err := os.MkdirAll(filepath.Dir(destPath), 0777)
	panicErr(err, "could not mkdir")

	dest, err := os.OpenFile(destPath, writeFlag, writePerm)
	panicErr(err, "could not open file for writing")
	defer dest.Close()

	src, err := os.Open(srcPath)
	panicErr(err, "could not open file for reading")
	defer src.Close()

	_, err = io.Copy(dest, src)
	panicErr(err, "could not copy files")
}

func panicErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
