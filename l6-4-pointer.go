// pointers

package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func changeName(u *User) {
	u.Name = "asdf"
	fmt.Println("New name: ", u.Name)
}

func checkPointer1() {
	u1 := User{"Jack", 23}
	fmt.Println("Before: ", u1.Name)
	// create pointer by taking the address of u1
	ptr := &u1
	changeName(ptr)
	fmt.Println("After: ", u1.Name)

	ptr_var := &User{"Jack", 28}
	fmt.Println("Before2: ", *ptr_var)
	changeName(ptr_var)
	fmt.Println("After2: ", *ptr_var)
}

func (u *User) Titlize() {
	u.Name = strings.ToTitle(u.Name)
}

func (u User) Reset() {
	u.Name = ""
	u.Age = 0
}

func checkPointer2() {
	u1 := User{"Mark", 67}
	fmt.Println("Before: ", u1)
	u1.Titlize()
	fmt.Println("After Titlize: ", u1)
	u1.Reset()
	// change by Reset is not reflected as it works on copy
	fmt.Println("After Reset: ", u1)
}

func checkPointer3() {
	s := new(string)
	*s = "hello"

	i := new(int)
	*i = 23

	u1 := new(User)
	u2 := &User{}

	fmt.Printf("s: %v, *s: %q\n", s, *s)
	fmt.Printf("i: %v, *i: %q\n", i, *i)
	fmt.Printf("u1: %+v, *u1: %+v\n", u1, *u1)
	fmt.Printf("u2: %+v, *u2: %+v\n", u2, *u2)
	
}

func main() {
	checkPointer1()
	checkPointer2()
	checkPointer3()
}
