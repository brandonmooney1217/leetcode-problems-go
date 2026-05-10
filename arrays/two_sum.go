package arrays

// LeetCode #1 (Easy): https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}

		seen[num] = i
	}
	return nil
}
