package _2_add_two_numbers

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
