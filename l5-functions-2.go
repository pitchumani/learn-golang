package main

import "fmt"

// variadic arguments
// it should be final argument
func sayHello(names ...string) {
	for _, n := range names {
		fmt.Println("Hello", n)
	}
}

func checkVarArgFunc() {
	sayHello("joe", "jack", "biden")
	sayHello("Mark")
	sayHello()

	// list can't be passed directly for variadic arguments
	users := []string{"Mani", "Sash", "Kesav", "Benit"}
	//sayHello(users)  // error
	sayHello(users...) // use variadic operator to expand the list
}

// slice type parameter
func LookupUsers(ids []int) {
	fmt.Println("looking up ids: ", ids)
}

// change above slice parameter into var arg func
func LookupUsers1(ids ...int) {
	fmt.Println("looking up ids: ", ids)
}

func main() {
	// pass slices
	checkVarArgFunc()

	// passing slice to slice argument
	id1 := 1
	id2 := 2
	id3 := 3
	ids := []int{id1, id2, id3}
	LookupUsers([]int{id1})
	LookupUsers(ids)

	// passing values to variadic argument
	LookupUsers1(id1, id2)
	LookupUsers1(ids...)
}
