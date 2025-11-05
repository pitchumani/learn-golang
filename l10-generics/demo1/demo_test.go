package demo

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Keys(t *testing.T) {
	t.Parallel()

	m := map[MyInt]string {
		1: "one",
		2: "two",
		3: "three",
	}

	act := Keys(m)

	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})

	exp := []float64{1, 2, 3}

	al := len(act)
	el := len(exp)
	if al != el {
		t.Fatalf("expected len %d, but got %d", el, al)
	}

	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)

	if at != et {
		t.Fatalf("expected type %s, but got %s", et, at)
	}
}

func Test_KEYS(t *testing.T) {
	t.Parallel()

	m := map[string]int {
		"one":1,
		"two":2,
		"three":3,
	}

	act := KEYS(m)

	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})

	exp := []string{"one", "two", "three"}

	al := len(act)
	el := len(exp)
	if al != el {
		t.Fatalf("expected len %d, but got %d", el, al)
	}

	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)

	if at != et {
		t.Fatalf("expected type %s, but got %s", et, at)
	}
}
