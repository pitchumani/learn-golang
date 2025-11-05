package main

import (
	"context"
	//"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// set self cancel with timeout of 10ms
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Millisecond)
	defer cancel()

	print(ctx)
}

