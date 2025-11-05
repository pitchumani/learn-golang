package wrap

type Unwrapper interface {
	Unwrap() error
}

// create AsError interface
type AsError interface {
	As(target any) bool
}

// create IsError interface
type IsError interface {
	Is(target error) bool
}
