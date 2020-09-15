package main

import (
	"fmt"
	"io"
	"os"
)

func PrintHelloWorld(w io.Writer) error {
	_, err := w.Write([]byte("Hello, world!"))

	if err != nil {
		return fmt.Errorf("Failed to write hello world")
	}

	return nil
}

func main() {
	PrintHelloWorld(os.Stdout)
}
