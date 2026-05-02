package sliding_window

import "testing"

func TestMaximumUniqueSubarray(t *testing.T) {

	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{4, 2, 4, 5, 6}, 17},
	}

	for _, tt := range tests {
		got := maximumUniqueSubarray(tt.nums)
		if tt.expected != got {
			t.Errorf("maximumUniqueSubarray(%v) = %d, want %d", tt.nums, got, tt.expected)
		}
	}
}
