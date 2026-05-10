package stack

// Leetcode 1944: https://leetcode.com/problems/number-of-visible-people-in-a-queue/description/

func canSeePersonsCount(heights []int) []int {

	n := len(heights)
	res := make([]int, n)

	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		viewCount := 0
		for len(stack) > 0 && heights[i] > stack[len(stack)-1] {
			viewCount++
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			viewCount++
		}

		stack = append(stack, heights[i])
		res[i] = viewCount
	}

	return res

}
