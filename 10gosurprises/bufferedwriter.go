package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Start OMIT
func main() {
	var w io.Writer = bufio.NewWriter(os.Stdout) // HL
	SendContent(w, "Hello world")
}

func SendContent(w io.Writer, content string) {
	fmt.Fprintf(w, "%s\n", content)
}

// End OMIT
