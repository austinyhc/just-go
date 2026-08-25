package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		name   string
		height []int
		want   int
	}{
		{"leetcode example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"leetcode example 2", []int{1, 1}, 1},
		{"two elements, uneven", []int{4, 3}, 3},
		{"increasing heights", []int{1, 2, 3, 4, 5}, 6},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := MaxArea(c.height); got != c.want {
				t.Errorf("MaxArea(%v) = %d, want %d", c.height, got, c.want)
			}
		})
	}
}
