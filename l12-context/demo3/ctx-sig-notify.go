package main

// context with timeout
// also checks for signal interrupt
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	fmt.Println("waiting for context to finish")
	<-ctx.Done()

	fmt.Println("context is finished")
}
