package dynamicprogramming

// LeetCode #70 (Easy): https://leetcode.com/problems/climbing-stairs/description/

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	s1, s2 := 1, 2

	for i := 3; i <= n; i++ {
		tmp := s2
		s2 = s2 + s1
		s1 = tmp
	}
	return s2

}
