package demo

import (
	"context"
	"math/rand/v2"
	"time"
	"testing"
)

func Test_Store(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)

	defer cancel()

	s := &Store{
		data: map[int]bool{},
	}

	const N = 10

	// goroutine sleeps randomnly
	go func() {
		for i := 0; i < N; i++ {
			r := rand.IntN(100)
			time.Sleep(time.Duration(r) * time.Millisecond)
			s.Add(i)
		}

		cancel()
	}()

	go func() {
		var i int
		for {
			r := rand.IntN(100)
			time.Sleep(time.Duration(r) * time.Millisecond)
			s.Get(i)
			i++
		}
	}()

	<-ctx.Done()

	act := len(s.data)
	if act != N {
		t.Fatalf("got %d, expected %d", act, N)
	}
}

