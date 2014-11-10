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
type Flusher interface { // HL
	Flush() error // HL
} // HL

func main() {
	var w io.Writer = bufio.NewWriter(os.Stdout)
	SendContent(w, "Hello world")
}

func SendContent(w io.Writer, content string) {
	fmt.Fprintf(w, "%s\n", content)
	f := w.(Flusher) // HL
	f.Flush()        // HL
}

// End OMIT
