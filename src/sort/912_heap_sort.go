package sort

import (
	"errors"
	"fmt"
)

func Swap(arr []int, i int, j int) error {
	if i >= len(arr) || j >= len(arr) {
		return errors.New("out of range")
	}
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
	return nil
}

func HeapSort(arr []int) {
	BuildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		swapIdx := 0
		Swap(arr, i, swapIdx)
		fmt.Printf("swap idx: %v, i:%v, after swap arr:%v\n", swapIdx, i, arr)
		adjIdx := AdjustDown(arr, swapIdx, i-1)
		for adjIdx != -1 {
			adjIdx = AdjustDown(arr, adjIdx, i-1)
		}
	}
}

func AdjustDown(arr []int, i int, endIdx int) int {
	left := i*2 + 1
	right := i*2 + 2
	biggest := arr[i]
	biggestIdx := i
	if left <= endIdx && arr[left] > biggest {
		biggest = arr[left]
		biggestIdx = left
	}
	if right <= endIdx && arr[right] > biggest {
		biggest = arr[right]
		biggestIdx = right
	}
	if i != biggestIdx {
		Swap(arr, i, biggestIdx)
		return biggestIdx
	}
	return -1
}

func BuildHeap(arr []int) {
	arrLen := len(arr)
	if len(arr) <= 1 {
		return
	}

	for i := (arrLen - 1) / 2; i >= 0; i-- {
		adjIdx := AdjustDown(arr, i, arrLen-1)
		for adjIdx != -1 {
			adjIdx = AdjustDown(arr, adjIdx, arrLen-1)
		}
	}
}
