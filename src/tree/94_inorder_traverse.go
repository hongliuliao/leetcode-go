package tree

import (
	"fmt"

	"github.com/leetcode-go/src/stack"
)

func InOrderTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	stack := stack.Stack{}
	stack.Push(root)
	leftNode := root.Left
	for leftNode != nil {
		stack.Push(leftNode)
		leftNode = leftNode.Left
	}

	for stack.Peek() != nil {
		element, err := stack.Pop()
		if err != nil {
			fmt.Printf("pop element err:%v\n", err)
			continue
		}
		treeNode := element.(*TreeNode)
		result = append(result, treeNode.Val)
		if treeNode.Right != nil {
			stack.Push(treeNode.Right)
			leftNode := treeNode.Right.Left
			for leftNode != nil {
				stack.Push(leftNode)
				leftNode = treeNode.Left
			}
		}
	}
	return result
}
