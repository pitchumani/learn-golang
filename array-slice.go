package main

import "fmt"
import "strings"

func func1() {
	namesArray := [4]string{"Mani", "Sash", "John"}
	namesSlice := []string{"Mark", "Akram", "Yogi"}

	fmt.Printf("namesArray: %[1]T: %[1]v\n", namesArray)
	fmt.Printf("namesSlice: %[1]T: %[1]v\n", namesSlice)

	var a [5]int
	var b [4]string
	var c [3]bool
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
}

func func2() {
	a1 := [2]string{"one", "two"}
	var a2 [2]string

	a2 = a1
	fmt.Println(a2)
	// error
	//a3 := [3]string{}
	//a3 = a1
}

func func3() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	// arrays have different memory, only values copied on assignment
	a2 = a1
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)

	a1[0] = "three"

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
}

func func4() {
	var names []string

	// generic function to append and capture the result back to the variable names
	names = append(names, "Kurt")
	fmt.Println(names)

	names = append(names, "Janis", "Jimi")
	fmt.Println(names)

	more := []string{"mani", "sash"}
	//names = append(names, more) // it is an error -
	// an approach to append one slice to another
	for _, name := range more {
		names = append(names, name)
	}
	fmt.Println(names)
	// instead of loop, can use variadic operator (...)
	names = append(names, more...)
	fmt.Println(names)
}

func checkLenCap() {
	names := []string{"John", "Doe"}
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))
	names = append(names, "mani")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))
	names = append(names, "sash")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))
	names = append(names, "mark")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))

	// check growth rate of capacity in large sequences
	var s1 []int

	hat := cap(s1)
	for i := 0; i < 1_000_000; i++ {
		s1 = append(s1, i)
		c := cap(s1)
		// print the capacity whenever it changes
		if c != hat {
			fmt.Println(hat, c)
		}
		hat = c
	}
}

func checkMake() {
	a := []string{}
	var b []string
	c := make([]string, 0)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := make([]string, 2, 3)
	d = append(d, "mani", "sash")
	fmt.Printf("len: %d, cap: %d, values: %v\n", len(d), cap(d), d)	
}

func checkSubset() {
	names := []string{"mark", "jason", "joe"}
	fmt.Println(names)
	subset := names[:2]
	fmt.Println("subset(names[:2]):", subset)

	for i, g := range subset {
		subset[i] = strings.ToUpper(g)
	}
	// both subset and names modified
	fmt.Println("subset:", subset)
	fmt.Println("names:", names)

	// to create copy, use copy
	names1 := []string{"mark", "jason", "joe"}
	subset1 := make([]string, 3)
	copy(subset1, names1[:2])
	fmt.Println("subset1:", subset1)
	for i, g := range subset1 {
		subset1[i] = strings.ToUpper(g)
	}
	fmt.Println("subset1:", subset1)
	fmt.Println("names1:", names1)
	
	subset1 = append(subset, "akram")
	fmt.Println("subset1:", subset1)
	fmt.Println("names1:", names1)
}

func slicesOnly(names []string) {
	for i, name := range names {
		fmt.Printf("%d. %v\n", i+1, name)
	}
}

func main() {
	// func1()
	// func2()
	// func3()
	// generic function append
	// func4()
	// other example generic functions for collections are: len, cap
	// checkLenCap()
	// checkMake()
	// checkSubset()
	
	names := [4]string{"mark", "jason", "matt", "andrew"}
	// error to pass array to list type argument
	// error to cast the array to list
	slicesOnly(names[:])


}
