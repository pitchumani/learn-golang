package main

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
)

// structs

type User struct {
	Name string
	Age int
}

func struct1() {
	u1 := User{}
	fmt.Printf("%+v\n", u1)
	// multi line initialization requires field name
	u2 := User{
		Name: "Jack",
		Age: 12,
	}
	fmt.Printf("%+v\n", u2)
	// single line initialization doesn't require field names	
	u3 := User{"Mark", 27}
	fmt.Printf("%+v\n", u3)
}

/*
// struct tag
type UserInfo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
*/

func struct2() {
	type User struct {
		ID int
		Name string
		Phone string
		Password string
	}
	u := User{
		Name: "Jack",
		ID:  23,
		Password: "helloworld",
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(u); err != nil {
		log.Fatal(err)
	}
	// struct tag is not specified, so the struct field names are taken
	// as json field names
	// {"ID":23,"Name":"Jack","Phone":"","Password":"helloworld"}

}

func struct3() {
	type User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Phone    string `json:"phone,omitempty"`
		Password string `json:"-"`
	}
	u := User{
		Name: "Jack",
		ID:  23,
		Password: "helloworld",
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(u); err != nil {
		log.Fatal(err)
	}
	// struct tags are specified for json, omitempty, ignore for password field etc
	// {"id":23,"name":"Jack"}

}

func main() {
	struct1()
	struct2()
	struct3()
}
