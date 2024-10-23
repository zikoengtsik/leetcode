package _2_add_two_numbers

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type testCase struct {
	input
	output
}

type input struct {
	l1 *ListNode
	l2 *ListNode
}

type output struct {
	expected *ListNode
}

func isSameNode(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func testAddTwoNumbers(t *testing.T, f func(l1, l2 *ListNode) *ListNode) {
	testCases := []testCase{
		{
			input{
				l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			},
			output{
				expected: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
			},
		},
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcName = funcName[strings.Index(funcName, "problems/"):]

	for _, tc := range testCases {
		actual := f(tc.input.l1, tc.input.l2)
		if !isSameNode(actual, tc.output.expected) {
			t.Errorf("Test [%s] failed: input(%v), expected(%v)!= actual(%v)",
				funcName, tc.input, tc.output.expected, actual)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testAddTwoNumbers(t, addTwoNumbers)
	testAddTwoNumbers(t, addTwoNumbersII)
}
