package _1_two_sum

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

func twoSumII(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
