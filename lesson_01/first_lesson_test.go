package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(2, 3)
	want := 5
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
