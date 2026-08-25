// Practice for lessons/0002-two-pointers.html — Container With Most Water (medium, opposite-ends).
// https://leetcode.com/problems/container-with-most-water/
package twopointers

// MaxArea returns the maximum area of water a container formed by two of
// the given heights can hold.
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
