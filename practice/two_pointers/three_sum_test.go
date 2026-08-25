package twopointers

import (
	"sort"
	"testing"
)

// normalize sorts each triplet, then sorts the list of triplets, so that
// triplet/element order doesn't matter for comparison.
func normalize(triplets [][]int) [][]int {
	out := make([][]int, len(triplets))
	for i, t := range triplets {
		cp := append([]int(nil), t...)
		sort.Ints(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return len(out[i]) < len(out[j])
	})
	return out
}

func equalTriplets(a, b [][]int) bool {
	a, b = normalize(a), normalize(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"leetcode example 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"leetcode example 2", []int{0, 1, 1}, [][]int{}},
		{"leetcode example 3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ThreeSum(c.nums)
			if !equalTriplets(got, c.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", c.nums, got, c.want)
			}
		})
	}
}
