package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	n := Fib(3)
	expected := 2
	if n != expected {
		fmt.Printf("n = %d,expected = %d", n, expected)
		t.Error("Not equal")
	}
}

func BenchmarkFib(b *testing.B) {
	for b.Loop() {
		Fib(30)
	}
}
