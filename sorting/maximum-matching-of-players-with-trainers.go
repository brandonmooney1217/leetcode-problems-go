package sorting

// LeetCode #2410 (Medium): https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

import (
	"slices"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {

	slices.SortFunc(players, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(trainers, func(a, b int) int {
		return a - b
	})

	i, j := 0, 0
	count := 0

	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}

	return count
}
