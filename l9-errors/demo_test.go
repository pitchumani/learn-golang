package demo

import "testing"

func Test_DoSomething(t *testing.T) {
	t.Parallel()

	err := DoSomething(0)
	if err != nil {
		t.Fatal(err)
	}

	err = DoSomething(1)
	if err != nil {
		t.Fatal("expected nil, got", err)
	}
}
