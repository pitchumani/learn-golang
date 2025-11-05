package demo

import (
	"fmt"
)

// make the named return to set return value while
// recovering from the panic
func DoSomething(input int) (err error) {
	// implement defer function to recover from panic
	defer func() {
		p := recover()
		if p == nil {
			// no panic was raised
			// return cleanly
			return
		}

		// assert value of p is an error
		if e, ok := p.(error); ok {
			// assign the recovered error
			// to the named return
			err = e
			return
		}

		// create new error from panic value
		err = fmt.Errorf("panic %T %s", p, p)
	}()
	
	switch input {
	case 0:
		return nil
	case 1:
		panic("one")
	}
	return nil
}

