package sort

import "errors"

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

}

func AdjustDown(arr []int, i int) int {
	arrLen := len(arr)
	left := i*2 + 1
	right := i*2 + 2
	biggest := arr[i]
	biggestIdx := i
	if left <= arrLen-1 && arr[left] > biggest {
		biggest = arr[left]
		biggestIdx = left
	}
	if right <= arrLen-1 && arr[right] > biggest {
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
		adjIdx := AdjustDown(arr, i)
		for adjIdx != -1 {
			adjIdx = AdjustDown(arr, adjIdx)
		}
	}
}
