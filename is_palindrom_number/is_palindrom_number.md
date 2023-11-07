# is palindrom number

## Description
Given an integer x, return true if x is a palindrome, and false otherwise.

### Examples

- Example 1:

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```
- Example 2:
```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

- Example 3:
```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```


### Constraints:
```
-2**31 <= x <= 2**31 - 1
```

## Flows:
### IsPalindromNumber
```
Checks if the number is negative, the number is not palindrom.
Check if x between 0 and 9. If true, return true as single digit numbers are palindromes.
Convert x int to a string using strconv.Itoa(x).
Initialize left as 0 and right as the length of the string minus 1.
Iterate while left is less than right.
Check if the characters at left and right positions in the string are not equal. If true, return false as the number is not a palindrome.
Increment left and decrement right.
If the loop completes without returning false, return true as the number is a palindrome.
```

### IsPalindromNumberCopilot
```
Check if x is less than 0. If true, return false.
Initialize reverse as 0 and origin as x.
Iterate until x becomes 0:
Multiply reverse by 10 and add the remainder of x divided by 10.
Divide x by 10.
Check if origin is equal to reverse. If true, return true. Otherwise, return false.
```

## Benchmarks:
```
goos: windows
goarch: amd64
pkg: github.com/slavakukuyev/leetcode-go/is_palindrom_number
cpu: 12th Gen Intel(R) Core(TM) i7-1255U

=== RUN   BenchmarkIsPalindromNumber
BenchmarkIsPalindromNumber
BenchmarkIsPalindromNumber-12
31410980                43.68 ns/op           24 B/op          1 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/is_palindrom_number 15.215s


=== RUN   BenchmarkIsPalindromCopilot
BenchmarkIsPalindromCopilot
BenchmarkIsPalindromCopilot-12
84893246                14.68 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/is_palindrom_number 1.304s
```

## LeetCode results:
### IsPalindromNumber
* Runtime Details
	- 8ms
	- Beats **77.17%** of users with Go

* Memory Details
	- 4.47MB
	- Beats **51.72%** of users with Go


### IsPalindromNumberCopilot
* Runtime Details
	- 29ms
	- Beats **5.18%** of users with Go

* Memory Details
	- 4.27MB
	- Beats **97.74%** of users with Go





### IsPalindromNumberCopilot
```
before loop reverse 0
before the loop x 12345678987654321

reverse 1
x 1234567898765432
reverse 12
x 123456789876543
reverse 123
x 12345678987654
reverse 1234
x 1234567898765
reverse 12345
x 123456789876
reverse 123456
x 12345678987
reverse 1234567
x 1234567898
reverse 12345678
x 123456789
reverse 123456789
x 12345678
reverse 1234567898
x 1234567
reverse 12345678987
x 123456
reverse 123456789876
x 12345
reverse 1234567898765
x 1234
reverse 12345678987654
x 123
reverse 123456789876543
x 12
reverse 1234567898765432
x 1
reverse 12345678987654321
x 0
true
```
