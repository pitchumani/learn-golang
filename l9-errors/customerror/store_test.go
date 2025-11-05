package customerror

import (
	"testing"
)

func Test_Store_All(t *testing.T) {
	t.Parallel()

	s := &Store{
		data: map[string][]Model{},
	}

	_, err := s.All("users")
	if err == nil {
		t.Fatal("expected error, not nil")
	}

	exp := "users"
	e, ok := err.(ErrTableNotFound)
	if !ok {
		t.Fatalf("expected ErrTableNotFound, got %T", err)
	}

	act := e.Table
	if act != exp {
		t.Fatalf("expected %q, got %q", exp, act)
	}
	
	if e.OccurredAt.IsZero() {
		t.Fatal("expected non-zero time")
	}
}
