package store

import "testing"

func Test_Store_Insert(t *testing.T) {
	t.Parallel()

	s := &Store[string, User]{
		data: map[string]User{},
	}
	exp := User{Email: "mark@example.com"}

	err := s.Insert(exp)
	if err != nil {
		t.Fatal(err)
	}

	act, err := s.Find(exp.Email)
	if err != nil {
		t.Fatal(err)
	}

	if exp.Email != act.Email {
		t.Fatalf("expected %v, got %v", exp, act)
	}
}
