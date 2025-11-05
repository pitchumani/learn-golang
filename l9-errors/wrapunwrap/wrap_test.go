package wrap

import (
	"testing"
	"errors"
)

func Test_Unwrap(t *testing.T) {
	t.Parallel()

	original := errors.New("original error")
	wrapped := Wrapper(original)

	/*
	unwrapped := errors.Unwrap(wrapped)
	if unwrapped != original {
		t.Fatalf("expected %v, got %v", original, unwrapped)
	}
	*/

	act := ErrorA{}
	// use As function instead of type assertion
	ok := errors.As(wrapped, &act)
	if !ok {
		t.Fatalf("expected %v to act as %v",
			wrapped, act)
	}

	if act.err == nil {
		t.Fatalf("expected non-nil, got nil")
	}
}
