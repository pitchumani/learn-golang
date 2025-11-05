package main

import (
	"context"
	"fmt"
)

func listener(ctx context.Context, i int) {
	fmt.Printf("listener %d is waiting\n", i)

	// block until context is cancelled
	<-ctx.Done()

	fmt.Printf("listener %d is exiting\n", i)
}
