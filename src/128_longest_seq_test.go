package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	num := LongestConsecutive(nums)
	assert.Equal(t, 4, num)
}
