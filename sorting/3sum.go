package sorting

// LeetCode #15 (Medium): https://leetcode.com/problems/3sum/

import (
	"slices" // Required for Go 1.21+
)

func threeSum(nums []int) [][]int {

	var matrix [][]int
	n := len(nums)
	slices.Sort(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		target := nums[i] * -1
		left, right := i+1, n-1

		for left < right {
			sm := nums[left] + nums[right]
			if sm == target {
				matrix = append(matrix, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sm > target {
				right--

			} else {
				left++
			}
		}
	}

	return matrix
}
