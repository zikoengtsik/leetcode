# [144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)

## Description

Given the `root` of a binary tree, return the preorder traversal of its nodes' values.

Example 1:

```
Input: root = [1,null,2,3]
Output: [1,2,3]
```

Example 2:

```
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [1,2,4,5,6,7,3,8,9]
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

- The number of nodes in the tree is in the range `[0, 100]`.
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
```
