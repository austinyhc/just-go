// Practice for lessons/0002-two-pointers.html — 3Sum (medium, "finish early" stretch problem).
// https://leetcode.com/problems/3sum/
package twopointers

import "sort"

// ThreeSum returns all unique triplets in nums that sum to zero. Order of
// triplets, and order within a triplet, does not matter.
func ThreeSum(nums []int) [][]int {
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	result := [][]int{}

	for i := 0; i < len(sorted)-2; i++ {
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}

		left, right := i+1, len(sorted)-1
		for left < right {
			sum := sorted[i] + sorted[left] + sorted[right]
			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				result = append(result, []int{sorted[i], sorted[left], sorted[right]})
				left++
				right--
				for left < right && sorted[left] == sorted[left-1] {
					left++
				}
				for left < right && sorted[right] == sorted[right+1] {
					right--
				}
			}
		}
	}

	return result
}
