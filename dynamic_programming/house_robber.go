package dynamicprogramming

// LeetCode #198 (Medium): https://leetcode.com/problems/house-robber/description/

func rob(nums []int) int {

	p1, p2 := 0, 0

	for _, val := range nums {

		tmp := p1
		p1 = max(p2+val, p1)
		p2 = tmp
	}
	return p1
}
