package main

import "fmt"
import "time"

func listener(i int, quit <-chan struct{}) {
	fmt.Printf("listener %d is waiting\n", i)

	<-quit

	fmt.Printf("listener %d is exiting\n", i)
}

func main() {
	quit := make(chan struct{})

	for i := 0; i < 5; i++ {
		go listener(i, quit)
	}

	time.Sleep(10 * time.Millisecond)

	fmt.Println("Closing the quit channel")

	close(quit)

	time.Sleep(50 * time.Millisecond)
}
