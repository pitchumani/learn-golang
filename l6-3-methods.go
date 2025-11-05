// methods

package main

import "fmt"

type User struct {
	Name string
	Age int
}

// attaching the method to User struct type
func (u User) String() string {
	return fmt.Sprintf("%s is %d", u.Name, u.Age)
}

// attaching the method to MyInt type (which is an alias for int)
type MyInt int
func (i MyInt) greet() {
	fmt.Printf("Hello I'm int %d\n", i)
}

func method1() {
	u := User{"Sash", 12}
	fmt.Println(u.String())

	i := MyInt(23)
	i.greet()

	type MyUser User
	u2 := User{"Jack", 27}
	fmt.Println(u2.String())
}

// declare a function type
type Greeter func() string

// function with argument type of a function
func sayHello(greeter Greeter) {
	fmt.Println(greeter())
}

func method2() {
	// call sayHello function with anonymous function as argument
	sayHello(func() string{
		return "Hello World!"
	})
}

func main() {
	method1()
	method2()
}

