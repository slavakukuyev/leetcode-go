package validparentheses

func IsValidParentheses(s string) bool {
	var m = map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	var slice []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			slice = append(slice, s[i])
		case '}', ')', ']':
			if len(slice) == 0 {
				return false
			}

			if slice[len(slice)-1] != m[s[i]] {
				return false
			}

			slice = slice[:len(slice)-1]

		}
	}

	return len(slice) == 0
}

func IsValidParenthesesCopilot(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != c-1 && stack[len(stack)-1] != c-2 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
