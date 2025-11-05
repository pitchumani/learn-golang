package main

import "fmt"

func main() {
	i := 10
	fmt.Println("i: ", i)
	i++
	fmt.Println("i: ", i)
	// below is an error - i++ is statement, not expression,
	// so, it cannot be used in expressions
	//fmt.Println("i: ", i++)
}
