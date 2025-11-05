package main

import "fmt"

func sayHello(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayHello2(greeting string, name string) {
	fmt.Println(greeting, name)
}

// function with multiple return values
func info(s []string)(string, int, int) {
	gs := fmt.Sprintf("%#v", s)
	l := len(s)
	c := cap(s)
	return gs, l, c
}

func f1()(retval bool) {
	retval = false
	fmt.Println("Function has named return")
	fmt.Println("return statement without value")
     return
}

func f2()(retval int) {
	fmt.Println("retval: ", retval)
	defer func() {
		// prints 45 as return statement updates the return value,
		// defer function is called after return
		fmt.Println("retval: ", retval)
		retval = 3
	}()
	return 45
}

func f4(fn func()) {
	fmt.Println("Function with function argument")
	fn()
}

func f5(fn func()) {
	fn()
}

func main() {
	sayHello("Hello", "Jason")
	sayHello2("Hi", "Akram")

	names := []string{"Jason", "Akram", "Joe"}
	
	gs, l, c := info(names)
	fmt.Println(gs)
	fmt.Println(l)
	fmt.Println(c)

	fmt.Println("Named return of f1: ", f1())

	f2();

	f3 := func() {
		fmt.Println("Function without return value")
	}

	f3()

	f4(f3)

	// anonymous function as argument
	n := "Janis"
	f5(func(){
		fmt.Println("hello ", n)
	})
}
