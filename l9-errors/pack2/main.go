package main

import (
	"bytes"
	"fmt"
	"io"
)

// check the input map for nil and throw error
// instead of panic on nil map
func AddUsers(m map[string]int) error {
	if m == nil {
		return fmt.Errorf("map is nil")
	}

	m["Amy"] = 27
	m["Mark"] = 68

	return nil
}

func Greet(w io.Writer) error {
	// invalid type assertion can cause panic
	//bb := w.(*bytes.Buffer)

	bb, ok := w.(*bytes.Buffer)
	if !ok {
		return fmt.Errorf("%T is not a bytes buffer", w)
	}
	_, err := bb.WriteString("Hello World")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// accessing uninitialized map will cause panic
	// var m map[string]int
	// so initialize it with :=
	m := map[string]int{}
	m["Amy"] = 27

	fmt.Printf("%+v\n", m)
}
