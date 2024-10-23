package _1_two_sum

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
	nums   []int
	target int
}

type output struct {
	expected []int
}

func testTwoSum(t *testing.T, f func(nums []int, target int) []int) {
	testCases := []testCase{
		{
			input{[]int{}, 1},
			output{nil},
		},
		{
			input{[]int{1, 2}, 2},
			output{nil},
		},
		{
			input{[]int{2, 7, 11, 15}, 9},
			output{[]int{0, 1}},
		},
		{
			input{[]int{3, 2, 4}, 6},
			output{[]int{1, 2}},
		},
		{
			input{[]int{3, 3}, 6},
			output{[]int{0, 1}},
		},
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcName = funcName[strings.Index(funcName, "problems/"):]
	for _, tc := range testCases {
		actual := f(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(actual, tc.output.expected) {
			t.Errorf("Test [%s] failed: input(%v), expected(%v)!= actual(%v)", funcName, tc.input, tc.output.expected, actual)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testTwoSum(t, twoSum)
	testTwoSum(t, twoSumII)
}
