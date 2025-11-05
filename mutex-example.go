package main

import (
	"fmt"
	"sync"
)

var (
    count int
    mu sync.Mutex
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}

func decrement() {
	mu.Lock()
	defer mu.Unlock()
	count--
}

func main() {
	fmt.Println("count:", count)
	go increment()
	go decrement()
	fmt.Println("count:", count)
}
