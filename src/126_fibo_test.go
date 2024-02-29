package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	assert.Equal(t, 4, Fibo(5), "Fibo assert err")
}
