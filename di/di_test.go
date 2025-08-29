package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Cheems")

	got := buffer.String()
	want := "Hello, Cheems"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
