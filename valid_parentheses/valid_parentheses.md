# Is Valid Parentheses

## Problem
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

## Examples:

* Example 1:
```
Input: s = "()"
Output: true
```

* Example 2:
```
Input: s = "()[]{}"
Output: true
```
* Example 3:
```
Input: s = "(]"
Output: false
```

## Constraints:
```
1 <= s.length <= 10**4
s consists of parentheses only '()[]{}'.
```

## Flows:
### IsValidParentheses
```
Initialize a stack.
Iterate over the characters in the string:
	Check if the character is an opening bracket. If true, push the character to the stack.
	Otherwise, check if the stack is empty or the top of the stack is not the corresponding opening bracket. If true, return false.
	Otherwise, pop the top of the stack.
Check if the stack is empty. If true, return true. Otherwise, return false.
```

### IsValidParenthesesCopilot
```
Initialize a stack.
Iterate over the characters in the string:
	Check if the character is an opening bracket. If true, push the character to the stack.
	Otherwise, check if the stack is empty or the top of the stack is not the corresponding opening bracket. If true, return false.
	Otherwise, pop the top of the stack.
Check if the stack is empty. If true, return true. Otherwise, return false.
```

## Benchmarks:
```
goos: windows
goarch: amd64
pkg: github.com/slavakukuyev/leetcode-go/valid_parentheses
cpu: 12th Gen Intel(R) Core(TM) i7-1255U

=== RUN   BenchmarkIsValidParentheses
BenchmarkIsValidParentheses
BenchmarkIsValidParentheses-12
 8428456               138.9 ns/op             8 B/op          1 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/valid_parentheses   1.356s


=== RUN   BenchmarkIsValidParenthesesCopilot
BenchmarkIsValidParenthesesCopilot
BenchmarkIsValidParenthesesCopilot-12
21976624                53.16 ns/op           24 B/op          2 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/valid_parentheses   2.023s

```