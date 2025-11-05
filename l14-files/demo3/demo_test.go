package demo

import (
	"strings"
	"testing"
)

func Test_Walk(t *testing.T) {
	t.Parallel()

	cab := createTestFS(t)
	
	exp := []string{
		"a.txt",
		"b.txt",
		"e/f/g.txt",
		"e/f/h.txt",
		"e/j.txt",
	}
	act, err := Walk(cab)
	if err != nil {
		t.Fatal(err)
	}

	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if as != es {
		t.Fatalf("expected %s, got %s", es, as)
	}
}
