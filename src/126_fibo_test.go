package main

import (
	"testing"
)

func TestFibo(t *testing.T) {
	result := Fibo(5)
	if result != 5 {
		t.Errorf("Fib(5) failed, expected %d but got %d", 5, result)
	}
}
