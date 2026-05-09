package binarysearch

// https://leetcode.com/problems/koko-eating-bananas/description/

func minEatingSpeed(piles []int, h int) int {

	left, right := 1, 0
	test_rate, test_hours := 0, 0
	for _, num := range piles {
		if num > right {
			right = num
		}
	}

	res := right

	for left <= right {
		test_rate = (right + left) / 2
		test_hours = 0
		for _, pile := range piles {
			tmp := (pile + test_rate - 1) / test_rate
			test_hours = test_hours + tmp
		}

		if test_hours <= h {
			right = test_rate - 1
			if test_rate < res {
				res = test_rate
			}
		} else {
			left = test_rate + 1
		}
	}

	return res
}
