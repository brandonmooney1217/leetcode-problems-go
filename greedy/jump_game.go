package greedy

// LeetCode #55 (Medium): https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
	furthest := nums[0]

	for i, val := range nums {
		if furthest < i {
			return false
		}

		furthest = max(furthest, i+val)
	}
	return true
}
