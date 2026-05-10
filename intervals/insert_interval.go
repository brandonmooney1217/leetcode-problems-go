package intervals

// LeetCode #57 (Medium): https://leetcode.com/problems/insert-interval/

func insert(intervals [][]int, newInterval []int) [][]int {

	i := 0
	n := len(intervals)
	res := [][]int{}

	/*
		Add all interval that do not overlap with new interval
	*/
	for i < n && newInterval[0] > intervals[i][1] {
		res = append(res, intervals[i])
		i++
	}

	/*Merge step*/
	for i < n && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	res = append(res, newInterval)

	/*Add remaining interval that do not overlap with new interval*/

	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}
