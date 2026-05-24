package dynamicprogramming

// LeetCode #53 (Medium): https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {

	curr := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		curr = max(curr+nums[i], nums[i])
		res = max(res, curr)
	}
	return res

}
