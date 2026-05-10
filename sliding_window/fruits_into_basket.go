package sliding_window

// LeetCode #904 (Medium): https://leetcode.com/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	res, left, count := 0, 0, 0
	seen := make(map[int]int)
	n := len(fruits)

	for r := 0; r < n; r++ {
		curr := fruits[r]
		if _, ok := seen[curr]; ok {
			seen[curr]++
		} else {
			seen[curr] = 1
		}
		count++

		for len(seen) > 2 {
			left_value := fruits[left]
			seen[left_value]--
			count--
			if seen[left_value] == 0 {
				delete(seen, left_value)
			}
			left++
		}
		if count > res {
			res = count
		}

	}
	return res
}
