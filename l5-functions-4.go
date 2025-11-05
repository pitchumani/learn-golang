package main

// check init function

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	// see that init is executed before the main
	fmt.Println("main")

	// multiple init functions can be created
	// they are executed in the order of their definitions

	// if multiple packages has init, they are executed
	// in the order of their imports
}

