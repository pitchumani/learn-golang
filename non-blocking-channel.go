package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// messages doesn't have any value
	// default case is in effect
	select {
	case msg := <- messages:
		fmt.Println("Received message: ", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"

	// msg cannot be sent to messages channel
	// because the channel doesn't have buffer and there is no receiver.
	// default case is in effect
	select {
	case messages <- msg:
		fmt.Println("Sent message")
	default:
		fmt.Println("No message sent")
	}

	// multi-way non-blocking select
	// attempts non-blocking receives on both messages and signals channels
	// as there is no message is sent, default is in effect
	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	case sig := <-signals:
		fmt.Println("Received signal: ", sig)
	default:
		fmt.Println("No activity")
	}
}
