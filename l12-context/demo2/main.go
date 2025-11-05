package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	
	ctx, cancel := context.WithCancel(ctx)
	// ensure the cancel function is called at least once
	// to avoid leaking resources
	defer cancel()

	// create 5 listeners
	for i := 0; i < 5; i++ {
		// launch listener in a goroutine
		go listener(ctx, i)
	}

	// allow the listeners to start
	time.Sleep(time.Millisecond * 500)

	fmt.Println("canceling the context")

	// cancel the context
	cancel()

	// allow the listeners to exit
	time.Sleep(time.Millisecond * 500)
}
