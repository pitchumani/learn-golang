package wrap

func Wrapper(original error) error {
	original = ErrorA{err: original}
	original = ErrorB{err: original}
	original = ErrorA{err: original}
	return original
}

func WrapperLong(original error) error {
	return ErrorC{
		err: ErrorB{
			err: ErrorA{
				err: original,
			},
		},
	}
}
