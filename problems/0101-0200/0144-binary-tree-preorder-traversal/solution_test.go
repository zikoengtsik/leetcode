package _144_binary_tree_preorder_traversal

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

func testPreorderTraversal(t *testing.T, f func(root *TreeNode) []int) {
	testCases := []testCase{
		{
			input{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			output{[]int{1, 2, 3}},
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

func TestPreorderTraversal(t *testing.T) {
	testPreorderTraversal(t, preorderTraversal)
	testPreorderTraversal(t, preorderTraversalII)
}
