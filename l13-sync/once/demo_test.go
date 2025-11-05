package demo

import (
	"fmt"
	"testing"
)

func Test_Builder(t *testing.T) {
	t.Parallel()

	b := &Builder{}

	for i := 0; i < 5; i++ {
		err := b.Build()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("builder built")

		if !b.Built {
			t.Fatal("expected builder to be built")
		}
	}
}

