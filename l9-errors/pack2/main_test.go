package main

import (
	"os"
	"testing"
)

func Test_AddUsers(t *testing.T) {
	t.Parallel()

	var m map[string]int

	err := AddUsers(m)

	if err != nil {
		t.Fatal(err)	
	}

	m1 := map[string]int{}
	err = AddUsers(m1)

	if err != nil {
		t.Fatal(err)	
	}
}

func Test_Greet(t *testing.T) {
	t.Parallel()

	err := Greet(os.Stdout)

	if err != nil {
		t.Fatal(err)
	}
}
