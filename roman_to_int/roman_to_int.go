package romantoint

import "regexp"

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
	l := len(s)
	if l < 1 || l > 15 {
		return false
	}

	for i := 0; i < l; i++ {
		if _, ok := roman_map[s[i]]; !ok {
			return false
		}
	}

	if l == 1 {
		return true
	}

	invalidRegex := regexp.MustCompile(`IIII|XXXX|CCCC|MMMM|VV|LL|DD|IL|IC|ID|IXL|IM|XD|XM|LC|LD|LM|DM`)
	return !invalidRegex.MatchString(s)
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

func RomanToInt2(s string) int {
	if !isValidRoman(s) {
		return 0
	}

	hm := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	cur, total, v := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		v = hm[s[i]]
		if cur > v {
			total -= v
		} else {
			total += v
			cur = v
		}
	}
	return total
}
