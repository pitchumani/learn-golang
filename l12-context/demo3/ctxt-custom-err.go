package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancelCause(ctx)

	err := ctx.Err()
	fmt.Println("Err: ", err)

	// create cancelable ctxt with custom error
	cause := context.Cause(ctx)
	fmt.Println("cause:", cause)

	cancel(fmt.Errorf("boom"))

	err = ctx.Err()
	fmt.Println("Err:", err)

	cause = context.Cause(ctx)
	fmt.Println("cause:", cause)
}
