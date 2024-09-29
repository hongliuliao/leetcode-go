package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	arr := []int{1, 2}
	Swap(arr, 0, 1)
	require.Equal(t, []int{2, 1}, arr)
}

func TestBuildHeap(t *testing.T) {
	arr := []int{3, 1, 2}
	BuildHeap(arr)
	require.Equal(t, []int{3, 1, 2}, arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{4, 5, 3, 1, 2}
	HeapSort(arr)
	require.Equal(t, []int{1, 2, 3, 4, 5}, arr)
}
