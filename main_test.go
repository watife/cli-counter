package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	got := count(b, false, false)

	exp := 4

	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	got := count(b, true, false)

	exp := 1

	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	got := count(b, false, true)

	exp := 19

	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
