package tree

import (
	"fmt"

	"github.com/leetcode-go/src/stack"
)

func Push(stack *stack.Stack, element *TreeNode) {
	stack.Push(element)
	fmt.Printf("put element to stack:%v\n", element.Val)
}

func PostOrderTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	stack := stack.Stack{}
	Push(&stack, root)
	leftNode := root.Left
	for leftNode != nil {
		Push(&stack, leftNode)
		leftNode = leftNode.Left
	}
	lastVistNode := root
	for stack.Peek() != nil {
		element := stack.Peek()
		treeNode := element.(*TreeNode)
		fmt.Printf("pop element:%v\n", treeNode.Val)

		if treeNode.Right != lastVistNode && treeNode.Right != nil {
			Push(&stack, treeNode.Right)
			leftNode := treeNode.Right.Left
			for leftNode != nil {
				Push(&stack, leftNode)
				leftNode = treeNode.Left
			}
		} else {
			element, err := stack.Pop()
			if err != nil {
				fmt.Printf("pop element err:%v\n", err)
				continue
			}
			treeNode := element.(*TreeNode)
			fmt.Printf("pop element:%v\n", treeNode.Val)
			lastVistNode = treeNode
			result = append(result, treeNode.Val)
		}

	}
	return result
}
