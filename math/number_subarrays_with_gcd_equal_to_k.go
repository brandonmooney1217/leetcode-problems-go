package math

// LeetCode #2447 (Medium): https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/description/
func subarrayGCD(nums []int, k int) int {
	length, res := len(nums), 0
	for i := 0; i < length; i++ {
		j := i
		currGcd := getGcd(nums[i], nums[j])

		for j < length {
			currGcd = getGcd(currGcd, nums[j])
			if currGcd == k {
				res++
			}
			j++

			if currGcd < k {
				break
			}
		}
	}
	return res
}

func getGcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
