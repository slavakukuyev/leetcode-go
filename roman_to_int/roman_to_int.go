package romantoint

import "strings"

var roman_map = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func isValidRoman(s string) bool {
	if len(s) < 1 || len(s) > 15 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if _, ok := roman_map[s[i]]; !ok {
			return false
		}
	}

	if len(s) == 1 {
		return true
	}

	if strings.Contains(s, "IIII") ||
		strings.Contains(s, "XXXX") ||
		strings.Contains(s, "CCCC") ||
		strings.Contains(s, "MMMM") ||
		strings.Contains(s, "VV") ||
		strings.Contains(s, "LL") ||
		strings.Contains(s, "DD") {
		return false
	}

	return true
}

func RomanToInt(s string) int {
	if !isValidRoman(s) {
		return 0
	}

	l := len(s) - 1
	end := false

	var number int
	for i := 0; i < l; i++ {

		if (s[i] == 'I' && s[i+1] == 'V') ||
			(s[i] == 'I' && s[i+1] == 'X') ||
			(s[i] == 'X' && s[i+1] == 'L') ||
			(s[i] == 'X' && s[i+1] == 'C') ||
			(s[i] == 'C' && s[i+1] == 'D') ||
			(s[i] == 'C' && s[i+1] == 'M') {
			number += (roman_map[s[i+1]] - roman_map[s[i]])
			i++
			if i == l {
				end = true
			}
			continue
		}

		number += roman_map[s[i]]

	}
	if !end {
		number += roman_map[s[len(s)-1]]
	}

	return number
}
