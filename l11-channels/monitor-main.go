package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

/*
// not a graceful shutdown when interrupt (Ctrl+c)
func main() {
	quit := make(chan struct{})

	mon := Monitor{}

	// start the gorooutine (thread)
	go mon.Start(quit)

	time.Sleep(10 * time.Second)

	close(quit)
}
*/

// capture signal and shutdown
func main() {
	// create a channel of os signal type
	sig := make(chan os.Signal)

	// notify sig channel in case of interrupt
	signal.Notify(sig, os.Interrupt)

	// create a channel of struct type
	quit := make(chan struct{})

	// create the instance of Monitor struct
	mon := Monitor{
		done: make(chan struct{}),
	}

	// start the thread to listen for interrupt
	go mon.Start(quit)

	// read the signal
	<-sig

	// close the quit channel
	close(quit)

	// block and wait for done channel
	// this may sometime cause indefinite wait
	//<-mon.Done()

	select {
	case <-mon.Done():
		os.Exit(0)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timed out while trying to shut down the monitor")
		os.Exit(1)
	}
}
