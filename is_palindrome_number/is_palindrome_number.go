package ispalindromenumber

import "strconv"

func IsPalindromeNumber(x int) bool {
	if x < 0 || x%10 == 0 || x == 0 {
		return false
	}

	if x > 0 && x <= 9 {
		return true
	}

	str := strconv.Itoa(x)
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func IsPalindromeNumberCopilot(x int) bool {
	if x < 0 {
		return false
	}

	var (
		reverse int
		origin  = x
	)

	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return origin == reverse
}
