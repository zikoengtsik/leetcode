# [1. Two Sum](https://leetcode.com/problems/two-sum)

## Description

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Example 2:

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

Example 3:

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- $2$ <= nums.length <= $10^{4}$
- $-10^{9}$ <= nums[i] <= $10^{9}$
- $-10^{9}$ <= target <= $10^{9}$
- Only one valid answer exists.

## Solutions

### I Brute Force

Make two nested loop to traverse the array `nums`. If `target - nums[i]` is in the array, return the index of `nums[i]` and the index of `target - nums[i]`.

- Time complexity: $O(n^{2})$
- Space complexity: $O(1)$

```go
func twoSum(nums []int, target int) []int {
  for i, num := range nums {
		j := i + 1
		for _, diff := range nums[j:] {
			if diff == target-num {
				return []int{i, j}
			}
			j += 1
		}
	}
	return nil
}
```

### II Hash Table

Traverse the array `nums`. If `target - nums[i]` is in the hash table, return the index of `nums[i]` and the index of `target - nums[i]`. Otherwise, store `nums[i]` and its index in `numMap`.

- Time complexity: $O(n)$
- Space complexity: $O(n)$

```go
func twoSumIII(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
```
