package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed successfully!")
		return nil
	case <-ctx.Done():
		fmt.Println("Operation cancelled!")
		return ctx.Err() // Return the cancellation error
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure cancel is called to release resources

	err := longRunningOperation(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
