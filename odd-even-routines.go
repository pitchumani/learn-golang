package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	cv *sync.Cond
	counter int
	limit int
)

func printOdd() {
	for {
		mu.Lock()
		for (counter % 2) == 0 && counter < limit {
			cv.Wait()
		}
		if (counter >= limit) {
			mu.Unlock()
			break
		}
		fmt.Printf(" o%v", counter)
		counter++
		cv.Signal()
		mu.Unlock()
	}
}

func printEven() {
	for {
		mu.Lock()
		for (counter % 2) == 1 && counter < limit {
			cv.Wait()
		}
		if (counter >= limit) {
			mu.Unlock()
			break
		}
		fmt.Printf(" e%v", counter)
		counter++
		cv.Signal()
		mu.Unlock()
	}
}

func main() {
	counter = 1
	limit = 20
	// init condition
	cv = sync.NewCond(&mu)

	// create a wait group to sync go routines
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		printOdd()
	}()
	go func() {
		defer wg.Done()
		printEven()
	}()
	// wait for go routines to finish
	wg.Wait()
	fmt.Println("\nEnd\n")
}
