package sliding_window

// LeetCode #2962 (Medium): https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

func countSubarrays(nums []int, k int) int64 {

	res, left, length := int64(0), 0, len(nums)
	maxInt := 0

	for _, num := range nums {
		if num > maxInt {
			maxInt = num
		}
	}

	count := 0
	for right, val := range nums {
		if val == maxInt {
			count++
		}

		for count >= k {
			res += int64(length - right)
			if nums[left] == maxInt {
				count--
			}
			left++
		}
	}
	return res
}
