package main

import "fmt"

func listener(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("listener exit")
}

func main() {
	ch := make(chan int, 10)
	go listender(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	time.Sleep(20 * time.Millisecond)

}
