package main

import "fmt"

func main() {
	// accessing nil will cause SIGSEGV
	//a := &Admin{}
	a := &Admin{
		User: &User{name: "Andy"},
	}
	fmt.Println(a.String())
}
