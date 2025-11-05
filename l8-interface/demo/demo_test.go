package main

import (
	"bytes"
	"testing"
)

func Test_WriteNow(t *testing.T) {
	t.Parallel()

	s := "Hello World!"

	bb := &bytes.Buffer{}
	err := WriteNow(bb, s)

	if err != nil {
		t.Fatal(err)
	}

	act := bb.String()
	if act != s {
		t.Fatalf("expected %q, got %q", s, act)
	}
}

