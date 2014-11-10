package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Start Interface OMIT
type fileSystem interface {
	OpenFile(name string, flag int, perm os.FileMode) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Writer
	io.Closer
}

// End Interface OMIT

// Start osFS OMIT
// OS Filesystem
type osFS struct{}

func (osFS) OpenFile(name string, flag int, perm os.FileMode) (file, error) {
	return os.OpenFile(name, flag, perm)
}

func (osFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

var fs fileSystem = osFS{}

// End osFS OMIT

func main() {
	f, err := fs.OpenFile("myfile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "File content")
}
