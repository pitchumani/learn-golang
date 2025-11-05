package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines completed.")
}
