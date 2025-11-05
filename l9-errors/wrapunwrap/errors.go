package wrap

import "fmt"
import "errors"

type ErrorA struct {
	err error
}

func (e ErrorA) Error() string {
	return fmt.Sprintf("[ErrorA] %s", e.err)
}

// implement the Unwrap interface for ErrorA
func (e ErrorA) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}
	return e.err
}

// implement AsError interface for ErrorA
func (e ErrorA) As(target any) bool {
	ex, ok := target.(*ErrorA)
	if !ok {
		return errors.As(e.err, target)
	}
	(*ex) = e
	return true
}

// implement IsError interface for ErrorA
func (e ErrorA) Is(target error) bool {
	if _, ok := target.(ErrorA); ok {
		return true
	}
	return errors.Is(e.err, target)
}

type ErrorB struct {
	err error
}

// implement the Unwrap interface for ErrorA
func (e ErrorB) Unwrap() error {
	return errors.Unwrap(e.err)
}

func (e ErrorB) Error() string {
	return fmt.Sprintf("[ErrorB] %s", e.err)
}

type ErrorC struct {
	err error
}

// implement the Unwrap interface for ErrorA
func (e ErrorC) Unwrap() error {
	return errors.Unwrap(e.err)
}

func (e ErrorC) Error() string {
	return fmt.Sprintf("[ErrorC] %s", e.err)
}
