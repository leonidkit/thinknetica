package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	want := 55
	got := Fibonacci(10)

	if got[10] != want {
		t.Fatalf("expectation: %d\nreality: %d", want, got)
	}
}
