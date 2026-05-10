package dynamicprogramming

// LeetCode #2770 (Medium): https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/

func maximumJumps(nums []int, target int) int {

	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			res[i] = 0
		} else {
			res[i] = -1
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			difference := nums[j] - nums[i]

			if difference >= (-1*target) && difference <= target {
				if res[j] != -1 {
					res[i] = max(res[i], res[j]+1)
				}
			}
		}
	}

	return res[0]

}
