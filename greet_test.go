package main

import (
	"bytes"
	"testing"
)


func TestGreets(t *testing.T) {
    buffer := bytes.Buffer{}
    Greets(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}