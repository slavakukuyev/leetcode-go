package maxarea

func MaxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	w := 0
	h := 0

	for left < right {
		w = right - left
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		if w*h > max {
			max = w * h
		}
	}

	return max
}

func MaxAreaCopilot(height []int) int {

	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
