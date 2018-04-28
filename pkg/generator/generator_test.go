package generator

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	expected := []string{"0", "1", "1", "2", "3", "5"}

	fib := Fibonacci()

	for i := range expected {
		n := fib().String()

		if n != expected[i] {
			t.Errorf("fibonacci: got %s, expected %s", n, expected[i])
		}
	}
}
