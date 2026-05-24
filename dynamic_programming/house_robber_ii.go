package dynamicprogramming

// LeetCode #198 (Medium): https://leetcode.com/problems/house-robber-ii/

func rob2(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	var helper func(nums []int) int
	helper = func(nums []int) int {
		p1, p2 := 0, 0
		for _, val := range nums {
			tmp := p1

			p1 = max(p2+val, p1)
			p2 = tmp
		}
		return p1
	}
	return max(helper(nums[1:len(nums)]), helper(nums[0:len(nums)-1]))
}
