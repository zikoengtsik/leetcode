package _144_binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecursion(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	preorderTraversalRecursion(root.Left, values)
	preorderTraversalRecursion(root.Right, values)
}

// Preorder traversal: root -> left -> right
func preorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	preorderTraversalRecursion(root, &values)
	return values
}

// Preorder traversal: root -> left -> right
func preorderTraversalII(root *TreeNode) []int {
	values, stack, node := make([]int, 0), make([]*TreeNode, 0), root

	for node != nil || len(stack) > 0 {
		for node != nil {
			values = append(values, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			node = node.Left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return values
}
