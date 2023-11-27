package removeduplicates

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l < 3 {
		return l
	}

	max := false
	left := 0
	pointer := 1
	for pointer < l {
		if nums[left] == nums[pointer] && !max {
			max = true
			left++
		} else if nums[left] != nums[pointer] {
			max = false
			left++
		}

		nums[left] = nums[pointer]
		pointer++
	}

	return left + 1
}
