package two_pointer

/*
You are given an integer array height of length n. There are n vertical lines drawn such that
the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
*/
func maxWater(array []int) int {

	if array == nil {
		return 0
	}

	if len(array) == 0 || len(array) == 1 {
		return 0
	}

	maxArea := 0
	left := 0
	right := len(array) - 1
	for i := 0; left < right; i++ {
		area := min(array[left], array[right]) * (right - left)
		maxArea = max(area, maxArea)
		if array[left] <= array[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
