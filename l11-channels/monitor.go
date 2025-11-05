package main

import (
	"fmt"
	"time"
)

//type Monitor struct{}

type Monitor struct {
	done chan struct{}
}

func (m Monitor) Done() <-chan struct{} {
	return m.done
}

func (m Monitor) Start(quit chan struct{}) {
	// create a ticker for every 10 millisecond
	tick := time.NewTicker(10 * time.Millisecond)
	// stop the ticker when exiting this function
	defer tick.Stop()

	// infinitely check for quit channel
	for {
		select {
		case <-quit:
			fmt.Println("shutting down monitor")
			return
		case <-tick.C:
			fmt.Println("monitor check")
		}
	}
}
