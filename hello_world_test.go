package main

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Test if hello world is printed", func(t *testing.T) {
		var got bytes.Buffer

		PrintHelloWorld(&got)
		want := "Hello, world!"

		if got.String() != want {
			t.Errorf("We got: %s but wanted %s", got.String(), want)
		}
	})

}
