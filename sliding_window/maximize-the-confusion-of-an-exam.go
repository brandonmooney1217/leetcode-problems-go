package sliding_window

//LeetCode #2024 (Medium): https://leetcode.com/problems/maximize-the-confusion-of-an-exam/

func maxConsecutiveAnswers(answerKey string, k int) int {

	l, res := 0, 0
	trueCount, falseCount := 0, 0

	for r, val := range answerKey {
		if val == 'T' {
			trueCount++
		} else {
			falseCount++
		}

		maxCount := max(trueCount, falseCount)

		for (r - l + 1) > maxCount+k {
			if answerKey[l] == 'T' {
				trueCount--
			} else {
				falseCount--
			}

			maxCount = max(falseCount, trueCount)
			l++
		}
		res = max(res, r-l+1)
	}

	return res
}
