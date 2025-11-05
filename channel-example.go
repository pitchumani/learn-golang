package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        // Process the job
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

	fmt.Println("Creating three goroutines to handle jobs")
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

	fmt.Println("Adding 5 jobs to jobs channel")
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
	fmt.Println("Closed jobs channel")

    for a := 1; a <= 5; a++ {
        <-results
    }
	close(results)
	fmt.Println("Received all results")
}
