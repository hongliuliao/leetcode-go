package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// go.exe test -timeout 30s -run ^TestInOrderTraverse$ github.com/leetcode-go/src/tree --count=1 -v
func TestInOrderTraverse(t *testing.T) {
	tree := CreateTree(1, 2, 3, 4, 5)
	ret := InOrderTraverse(tree)
	for _, val := range ret {
		fmt.Println(val)
	}
	require.NotNil(t, ret)
}
