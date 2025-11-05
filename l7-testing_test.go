package demo

import "testing"

func Test_IsOdd(t *testing.T) {
	if (!IsOdd(1)) {
		t.Fatal("1 should be odd")
	}
	if (IsOdd(2)) {
		t.Fatal("2 should not be odd")
	}
}
