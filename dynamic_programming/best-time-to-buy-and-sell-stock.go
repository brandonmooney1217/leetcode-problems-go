package dynamicprogramming

// LeetCode #121 (Easy): https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {

	res := 0
	buyPrice := prices[0]

	for i := 0; i < len(prices); i++ {
		currPrice := prices[i]

		if currPrice < buyPrice {
			buyPrice = currPrice
		} else {
			res = max(res, currPrice-buyPrice)
		}
	}
	return res

}
