package main

import "fmt"

type Poet struct {
	Name string
}

func (p Poet) Perform() {
	fmt.Println(p.Name, "is reading poetry")
}

