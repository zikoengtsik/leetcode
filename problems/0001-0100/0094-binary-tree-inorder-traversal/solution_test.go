package _94_binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

type testCase struct {
	input
	output
}

type input struct {
	root *TreeNode
}

type output struct {
	expected []int
}

func testInorderTraversal(t *testing.T, f func(root *TreeNode) []int) {
	testCases := []testCase{
		{
			input{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			output{[]int{1, 3, 2}},
		},
		{
			input{nil},
			output{[]int{}},
		},
		{
			input{&TreeNode{Val: 1}},
			output{[]int{1}},
		},
	}

	for _, tc := range testCases {
		actual := f(tc.input.root)
		if !reflect.DeepEqual(actual, tc.output.expected) {
			t.Errorf("Test failed: input(%v), expected(%v)!= actual(%v)",
				tc.input, tc.output.expected, actual,
			)
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	testInorderTraversal(t, inorderTraversal)
	testInorderTraversal(t, inorderTraversalII)
}
