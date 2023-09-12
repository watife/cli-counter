package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	got := count(b)

	exp := 4

	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
