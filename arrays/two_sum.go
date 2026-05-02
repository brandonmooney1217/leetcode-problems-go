package arrays

/*
Problem: Two Sum (LeetCode #1)
Difficulty: Easy
Link: https://leetcode.com/problems/two-sum/


*/

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
