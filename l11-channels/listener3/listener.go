package main

import "fmt"

type User struct {
	ID int
	Name string
}

func main() {
	ch := make(chan User)

	go func() {
		ch <- User{ID: 1, Name: "mark"}
	}()

	user := <- ch

	fmt.Printf("read successful: %+v\n", user)
	close(ch)

	user, ok := <-ch

	if !ok {
		fmt.Println("attempted read of closed channel")
		fmt.Printf("received zero value %s\n", user)
		return
	}

	fmt.Printf("read successful: %+v\n", user)
}
