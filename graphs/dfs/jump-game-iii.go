package dfs

// LeetCode #1306 (Medium): https://leetcode.com/problems/jump-game-iii/

func canReach(arr []int, start int) bool {

	// check out of bounds
	if start < 0 || start >= len(arr) {
		return false
	}

	// check if val == 0
	if arr[start] == 0 {
		return true
	}

	// check if visited
	if arr[start] < 0 {
		return false
	}

	jumpValue := arr[start]
	arr[start] = arr[start] * -1

	canReachFromHere := canReach(arr, start+jumpValue) || canReach(arr, start-jumpValue)

	return canReachFromHere
}
