package main

import "fmt"

// type declaration

type MyInt int
type MyString string
type MyMap map[string]string

func main() {
	var i1 MyInt = 1
	println(i1)

	var m1 MyMap = map[string]string{"a": "A"}
	fmt.Println(m1)
	m2 := MyMap{"b": "B"}
	fmt.Println(m2)
}
