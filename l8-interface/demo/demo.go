package main

import (
	"bytes"
	"fmt"
	"io"
)
/*
func WriteNow(i any, s string) error {
	switch i.(type) {
	case *bytes.Buffer:
		fmt.Println("type was a *bytes.Buffer", s)
	case io.StringWriter:
		fmt.Println("type was a io.StringBuffer", s)
	case io.Writer:
		fmt.Println("type was a io.Writer", s)
	}

	return fmt.Errorf("can not write to %T", i)
}
*/

func WriteNow(i any, s string) error {
	switch t := i.(type) {
	case *bytes.Buffer:
		t.WriteString(s)
		return nil
	case io.StringWriter:
		t.WriteString(s)
		return nil
	case io.Writer:
		t.Write([]byte(s))
		return nil
	}

	return fmt.Errorf("can not write to %T", i)
}

