package _145_binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Postorder traversal: left -> right -> root
func postorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	postorderTraversalRecursion(root, &values)
	return values
}

func postorderTraversalRecursion(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	postorderTraversalRecursion(root.Left, values)
	postorderTraversalRecursion(root.Right, values)
	*values = append(*values, root.Val)
}

// Postorder traversal: left -> right -> root
func postorderTraversalII(root *TreeNode) []int {
	values, stack, node := make([]int, 0), make([]*TreeNode, 0), root

	for node != nil || len(stack) > 0 {
		for node != nil {
			values = append([]int{node.Val}, values...)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			node = node.Right
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return values
}
