package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	queue := make([]*TreeNode, 0)
	root := &TreeNode{Val: nums[0]}
	queue = append(queue, root)

	i := 1
	for len(queue) > 0 && i < len(nums) {
		current := queue[0]
		queue = queue[1:]

		if nums[i] != -1 {
			leftNode := &TreeNode{Val: nums[i]}
			current.Left = leftNode
			queue = append(queue, leftNode)
		}
		i++
		if i < len(nums) && nums[i] != -1 {
			rightNode := &TreeNode{Val: nums[i]}
			current.Right = rightNode
			queue = append(queue, rightNode)
		}
		i++
	}

	return root
}
