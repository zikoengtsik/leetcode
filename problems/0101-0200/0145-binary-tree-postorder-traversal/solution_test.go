package _145_binary_tree_postorder_traversal

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

func testPostorderTraversal(t *testing.T, f func(root *TreeNode) []int) {
	testCases := []testCase{
		{
			input{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			output{[]int{3, 2, 1}},
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

func TestPostorderTraversal(t *testing.T) {
	testPostorderTraversal(t, postorderTraversal)
	testPostorderTraversal(t, postorderTraversalII)
}
