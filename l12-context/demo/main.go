package main

import (
	"context"
	"demo/db"
	"demo/web"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = db.WithRequestID(ctx)

	ctx = web.WithRequestID(ctx)

	//id := ctx.Value(db.RequestID)
	id, err := db.RequestIDFrom(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db.RequestID: ", id)
}
