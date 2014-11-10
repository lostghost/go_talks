package main

import (
	"bytes"
	"os"
)

// Start OMIT
/*
testFS implements fileSystem using memory.

Test files are implemented as an in-memory map of filename to file.
*/
type testFS struct {
	f map[string]*testFile
}

/*
OpenFile is a replacement for the osFS OpenFile function. Rather
than creating a new file on disk, the new file will be created in
the files map in memory.
*/
func (fs *testFS) OpenFile(name string, flag int, perm os.FileMode) (file, error) {
	if fs.f == nil {
		fs.init()
	}
	f := &testFile{b: bytes.NewBuffer(nil)}
	fs.f[name] = f
	return f, nil
}

// End OMIT
