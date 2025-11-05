package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 2-value form of receive
			// second value will be false if jobs has been closed
			j, more := <- jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs. Send true to done channel.")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sent job: ", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs. Closed channel jobs")

	// receive from done channel
	<- done

	_, ok := <- jobs
	fmt.Println("Received more jobs: ", ok)
}
	
