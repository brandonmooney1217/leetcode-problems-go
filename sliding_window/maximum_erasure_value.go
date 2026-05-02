package sliding_window

// LeetCode #1695 (Medium): https://leetcode.com/problems/maximum-erasure-value/

func maximumUniqueSubarray(nums []int) int {
	left, res, sm := 0, 0, 0
	dct := make(map[int]int)

	for right := 0; right < len(nums); right++ {
		num := nums[right]
		sm += num
		dct[num]++

		for dct[num] >= 2 {
			dct[nums[left]]--
			sm -= nums[left]
			left++
		}
		if sm > res {
			res = sm
		}
	}
	return res
}
