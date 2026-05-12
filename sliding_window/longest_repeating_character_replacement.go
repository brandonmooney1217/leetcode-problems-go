package sliding_window

// LeetCode #424 (Medium): https://leetcode.com/problems/longest-repeating-character-replacement/

func characterReplacement(s string, k int) int {
	res, left, mx_char := 0, 0, 0
	dct := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		dct[s[right]]++
		if dct[s[right]] > mx_char {
			mx_char = dct[s[right]]
		}
		for (right-left+1)-mx_char > k {
			dct[s[left]]--
			left++
		}
		if (right - left + 1) > res {
			res = right - left + 1
		}
	}

	return res

}
