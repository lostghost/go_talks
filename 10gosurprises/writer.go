package main

import (
	"fmt"
	"io"
	"os"
)

// Start OMIT
func main() {
	var w io.Writer = os.Stdout
	SendContent(w, "Hello world")
}

func SendContent(w io.Writer, content string) {
	fmt.Fprintf(w, "%s\n", content)
}

// End OMIT
