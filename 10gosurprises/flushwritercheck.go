package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	_ = bufio.MaxScanTokenSize
)

// Start OMIT
type Flusher interface {
	Flush() error
}

func main() {
	var w io.Writer = bufio.NewWriter(os.Stdout)
	SendContent(w, "Hello world")
}

func SendContent(w io.Writer, content string) {
	fmt.Fprintf(w, "%s\n", content)
	if f, ok := w.(Flusher); ok { // HL
		f.Flush()
	}
}

// End OMIT
