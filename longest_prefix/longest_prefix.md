# Longest Prefix
## Problem
Given a list of strings, find the longest common prefix between all strings.

## Constraints:
```
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
```


## Flows:
#### LongestPrefix
1. Optional: The function first checks if the length of the input slice is less than 1 or greater than 200. If so, it returns an empty string.
2. Optional: It then checks if the length of the first string in the slice is less than 1 or greater than 200. If so, it also returns an empty string.
3. The function initializes a variable latest with the value of the first string in the slice.
4. It creates a buffer called prefix to store the common prefix.
5. It iterates over the remaining strings in the slice starting from the second string.
6. For each string, it resets the prefix buffer.
7. It then iterates over the characters of the latest string and the current string simultaneously.
8. If the characters at the same index are equal, it appends the character to the prefix buffer.
9. If the characters are not equal, it breaks out of the loop.
10. After iterating over all the strings, the function assigns the value of the prefix buffer to the latest variable.
11. Finally, it returns the latest string, which represents the longest common prefix among the input strings.

#### LongestPrefixLeeCode
1. Optional: The function first checks if the length of strs is 0. If it is, an empty string is returned.
2. The function then iterates over the characters of the first string in strs using a for loop.
3. For each character c at index i in the first string, the function checks if c is equal to the corresponding character at index i in the other strings in strs.
4. If c is not equal to the corresponding character in any of the other strings, the function returns the substring of the first string from index 0 to i.
5. If all characters in the first string match the corresponding characters in the other strings, the function returns the first string itself as the longest common prefix.

#### LongestPrefixCopilot
1. Optional: Check if the length of strs is 0. If it is, return an empty string.
2. Set the prefix variable to the first string in strs.
3. Iterate through the remaining strings in strs starting from the second string.
4. Check if the prefix is not a prefix of the current string. If it is not, remove the last character from the prefix.
5. If the prefix becomes empty, return an empty string.
6. Repeat steps 4 and 5 until the prefix is a prefix of all the strings or becomes empty.
7. Return the prefix.


## Example
```go 
Input: ["colorado", "color", "cold"]
Output: "col"
```

```go
Input: ["a", "b", "c"]
Output: ""
```

```go
Input: ["spot", "spotty", "spotted"]
Output: "spot"
```

## Complexity
Time: O(n) where n is the number of characters in the list of strings
Space: O(1)
```go
//Copilot example
func longestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
```

## Benchmarking
```shell
goos: windows
goarch: amd64
pkg: github.com/slavakukuyev/leetcode-go/longest_prefix
cpu: 12th Gen Intel(R) Core(TM) i7-1255U

=== RUN   BenchmarkLongestPrefix
BenchmarkLongestPrefix
BenchmarkLongestPrefix-12
 7937782               147.0 ns/op            84 B/op          9 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/longest_prefix      1.350s

=== RUN   BenchmarkLongestPrefixLeetCode
BenchmarkLongestPrefixLeetCode
BenchmarkLongestPrefixLeetCode-12
82867777                14.80 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/longest_prefix      1.273s

=== RUN   BenchmarkLongestPrefixCopilot
BenchmarkLongestPrefixCopilot
BenchmarkLongestPrefixCopilot-12
 6757080               163.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/longest_prefix      1.333s
```


## LeetCode Results

#### LongestPrefix
* Runtime Details
	- 0ms
	- Beats **100%** of users with Go

* Memory Details
	- 2.44MB
	- Beats **13.07%** of users with Go


#### LongestPrefixLeetCode
* Runtime Details
	- 2ms
	- Beats **72.82%** of users with Go

* Memory Details
	- 2.40MB
	- Beats **40.75%** of users with Go


#### LongestPrefixCopilot
* Runtime Details
	- 2ms
	- Beats **72.82%** of users with Go

* Memory Details
	- 2.38MB
	- Beats **40.75%** of users with Go