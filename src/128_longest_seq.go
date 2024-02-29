package main

func LongestConsecutive(nums []int) int {
	max := 0
	numCntSet := make(map[int]bool)

	// 使用 map 模拟 unordered_set
	for _, num := range nums {
		numCntSet[num] = true
	}

	for num := range numCntSet {
		left := num - 1
		right := num + 1
		seq := 1

		// 检查左侧是否存在连续序列的起点
		if !numCntSet[left] {
			// 递增检查右侧连续序列的长度
			for numCntSet[right] {
				delete(numCntSet, right) // 一个优化点, 减少无效计算
				seq++
				right++
			}
			max = Max(max, seq)
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
