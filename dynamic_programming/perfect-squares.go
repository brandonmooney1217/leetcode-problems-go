package dynamicprogramming

// LeetCode #279 (Medium): https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {

	var res []int
	for i := 1; (i * i) <= n; i++ {
		res = append(res, i*i)
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	for sm := 1; sm <= n; sm++ {
		for _, num := range res {
			if sm >= num {
				dp[sm] = min(dp[sm], 1+dp[sm-num])
			}
		}
	}
	return dp[n]
}
