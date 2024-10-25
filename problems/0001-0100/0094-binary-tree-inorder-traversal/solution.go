package _94_binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalRecursion(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorderTraversalRecursion(root.Left, values)
	*values = append(*values, root.Val)
	inorderTraversalRecursion(root.Right, values)
}

// In-order traversal: left -> root -> right
func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	inorderTraversalRecursion(root, &values)
	return values
}

// In-order traversal: left -> root -> right
func inorderTraversalII(root *TreeNode) []int {
	values, stack, node := make([]int, 0), make([]*TreeNode, 0), root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			values = append(values, node.Val)
			node = node.Right
		}
	}

	return values
}
