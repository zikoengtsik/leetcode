# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers)

## Description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

Example 2:

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

Example 3:

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

Constraints:

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`.
- It is guaranteed that the list represents a number that does not have leading zeros.

## Solutions

### I Iterative

Traverse two linked lists at the same time iteratively, and add the corresponding digits and carry.

- Time complexity: $O(max(m, n))$, where $m$ and $n$ are the lengths of the two linked lists.
- Space complexity: $O(1)$.

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, isCarry := ListNode{}, false
	p := &dummy

	for l1 != nil || l2 != nil || isCarry {
		val := 0
		if isCarry {
			val += 1
		}
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		isCarry = val > 9
		p.Next = &ListNode{Val: val % 10}
		p = p.Next
	}

	return dummy.Next
}
```

### II Recursive

Traverse two linked lists at the same time recursively, and add the corresponding digits and carry.

- Time complexity: $O(max(m, n))$, where $m$ and $n$ are the lengths of the two linked lists.
- Space complexity: $O(max(m, n))$, where $m$ and $n$ are the lengths of the two linked lists.

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersRecursion(l1 *ListNode, l2 *ListNode, isCarry bool) *ListNode {
	if l1 == nil && l2 == nil && !isCarry {
		return nil
	}
	node, val := ListNode{}, 0
	if isCarry {
		val += 1
	}
	if l1 != nil {
		val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		val += l2.Val
		l2 = l2.Next
	}

	isCarry = val > 9
	node.Val = val % 10
	node.Next = addTwoNumbersRecursion(l1, l2, isCarry)
	return &node
}

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRecursion(l1, l2, false)
}
```
