package main

import "fmt"

type Musician struct {
	Name string
}

func (m Musician) Perform() {
	fmt.Println(m.Name, "is singing")
}
