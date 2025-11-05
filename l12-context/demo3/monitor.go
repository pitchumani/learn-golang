package main

import (
	"context"
	"fmt"
	"time"
)

type Monitor struct {
	cancel context.CancelFunc
}

func (m *Monitor) Start(ctx context.Context) context.Context {
	go m.listen(ctx)

	ctx, cancel := context.WithCancel(context.Background())

	m.cancel = cancel

	return ctx
}

func (m *Monitor) listen(ctx context.Context) {
	defer m.cancel()

	tick := time.NewTicker(time.Millisecond * 10)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down monitor")

			m.cancel()

			return
		case <-tick.C:
			fmt.Println("monitor check")
		}
	}
}

