package dynamicprogramming

// LeetCode #746 (Easy): https://leetcode.com/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) int {
	down2, down1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		curr := min(down1, down2) + cost[i]
		down2 = down1
		down1 = curr
	}
	return min(down2, down1)
}
