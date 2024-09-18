package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	stack := Stack{}

	// 入栈几个元素
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")

	// 查看栈顶元素
	top := stack.Peek()
	if top == nil {
		fmt.Printf("Top element: %v\n", top)
	}
	require.NotNil(t, top)

	// 出栈几个元素
	for stack.Peek() != nil {
		element, err := stack.Pop()
		if err == nil {
			fmt.Printf("Popped element: %v\n", element)
		}
	}
}
