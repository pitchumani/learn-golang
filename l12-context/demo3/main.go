package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	// create a empty context
	ctx := context.Background()

	// wrapping the ctx with a cancellable context
	ctx, cancel := context.WithCancel(ctx)
	// cancel the context at the end
	defer cancel()

	// create Monitor instance
	mon := Monitor{}

	// start the monitor with the context
	ctx = mon.Start(ctx)

	// start a goroutine - that cancels after 50ms
	go func() {
		time.Sleep(time.Millisecond * 50)
		cancel()
	}()

	select {
	case <-ctx.Done():
		os.Exit(0)
	case <-time.After(time.Second * 2):
		fmt.Println("timed out while trying to shut down the monitor")

		if err := ctx.Err(); err != nil {
			fmt.Printf("error: %s\n", err)
		}

		os.Exit(1)
	}
}
