package tree

import (
	"fmt"

	"github.com/leetcode-go/src/stack"
)

func PostOrderTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	stack := stack.Stack{}
	stack.Push(root)
	isFirst := true
	for stack.Peek() != nil {
		element, err := stack.Pop()
		if err != nil {
			fmt.Printf("pop element err:%v\n", err)
			continue
		}
		treeNode := element.(*TreeNode)
		result = append(result, treeNode.Val)
		if isFirst {
			stack.Push(element)
			isFirst = false
		}

		if treeNode.Right != nil {
			stack.Push(treeNode.Right)
		}
		if treeNode.Left != nil {
			stack.Push(treeNode.Left)
		}
	}
	return result
}
