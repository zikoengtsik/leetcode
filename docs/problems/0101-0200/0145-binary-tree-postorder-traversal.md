# [145. Binary Tree Postorder Traversal]

## Description

Given the `root` of a binary tree, return the postorder traversal of its nodes' values.

Example 1:

```
Input: root = [1,null,2,3]
Output: [3,2,1]
```

Example 2:

```
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,6,7,5,2,9,8,3,1]
```

Example 3:

```
Input: root = []
Output: []
```

Example 4:

```
Input: root = [1]
Output: [1]
```

Constraints:

- The number of the nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`.

## Solutions

### I Recursion

Complexity:

- Time complexity: $O(n)$, where $n$ is the number of nodes in the tree.
- Space complexity: $O(h)$, where $h$ is the height of the tree.

```go
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
```

### II Iteration

Complexity:

- Time complexity: $O(n)$, where $n$ is the number of nodes in the tree.
- Space complexity: $O(n)$, where $n$ is the number of nodes in the tree.

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
```
