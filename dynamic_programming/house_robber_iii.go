package dynamicprogramming

// LeetCode #337 (Medium): https://leetcode.com/problems/house-robber-iii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	tmp := dfs(root)
	return max(tmp[0], tmp[1])
}

func dfs(node *TreeNode) [2]int {

	if node == nil {
		return [2]int{0, 0} // skip, rob
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	rob := node.Val + left[0] + right[0]
	skip := max(left[0], left[1]) + max(right[0], right[1])

	return [2]int{skip, rob}
}
