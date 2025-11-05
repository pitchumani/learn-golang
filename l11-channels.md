# Channels

- concurrency
- channels
- iteration and select statements
- closed channels
- buffered channels
- system signals

In Go, concurrency is built into the language. No need for any third party libraries,
system processes, or kernel threads.

## Concurrency

Go uses goroutine model to achieve concurrency.
goroutine
- independent function launched by go statement, capable of being run concurrently
with other goroutines.
- it is lightweight thread of execution that is managed by go runtime scheduler

func main() {
    go someFunction()

    fo func() {
        // do something
    }()
}

func someFunction() {
    // do something
}

Go Runtime Scheduler
It is responsible for distributing the runnable goroutines over multiple OS threads.
Work sharing and Work stealing modes

The runtime package provides a number of functions that can be used to query and
make small changes to the Go runtime.
e.g. GOMAXPROCS - to set number runtime threads that the Go runtime will use


Running the below program will not print anything. Because the program exits
before runtime scheduler has a chance to run its goroutines.

```go
package main

import "fmt"

func main() {
	go fmt.Println("Hello, World!")
}
```

## Channels
channels are a type conduit through which we can receive and send values.
characteristics
- typed
- synchronous - sender/receiver must wait before sending or receiving
- FIFO
- unbuffered or buffered - buffer channel will hold a limited number of values
  - when buffer is full, the sender will block until a value is received
  - when buffer is empty, the receiver will block until a value is sent
- directional - bi-directional or unidirectional
  - bi-directional - can send and receive values
  - unidirectional - can send or receive values
  
create a channel -- chan keyword followed by type, use make function
ch := make(chan int)  // ch is a channel that can send/receive int values

arrow operators to indicate the direction

## Iteration and Select statements


## buffered channels
like voice mails
can use a slice create a buffer to hold on the messages
writing into channel after the buffer is full, will lead to fatal error

## system signals
all programs should attempt a graceful shutdown
- detecting that the program was requested to shut down
- shut down all internal processes, including long running goroutines
- have a reasonable timeout in the event that internal processes are
  taking too long to shut down or are deadlocked
- respond to an actual user request for immediate hard shutdown
- record the result of shutdown(success, timeout, user intervention)


