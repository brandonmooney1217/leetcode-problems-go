package greedy

// Leetcode #1665 (Hard): https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/?envType=daily-question&envId=2026-05-12
import (
	"sort"
)

func minimumEffort(tasks [][]int) int {

	currEnergy, startEnergy := 0, 0
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
	})

	for i := 0; i < len(tasks); i++ {
		actual, min := tasks[i][0], tasks[i][1]

		if currEnergy < min {
			startEnergy += (min - currEnergy)
			currEnergy = min
		}
		currEnergy -= actual
	}

	return startEnergy
}
