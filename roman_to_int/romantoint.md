# roman to int

## Description
Given a roman numeral, convert it to an integer.

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.


Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.


### Examples
* Example 1:

```
Input: s = "III"
Output: 3
Explanation: III = 3.
```


* Example 2:

```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
```

* Example 3:

```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

### Validation contains:
1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].

## Solution explanation
### Solution 1
#### RomanToInt 

1. The function first checks if the input string needs to be validated. If validation is required and the input string is not a valid Roman numeral according to the isValidRoman function, the function returns 0.
2. The function initializes variables l (length of the input string - 1) and end (a flag indicating if the loop has reached the end of the input string).
3. The function then iterates over the input string from index 0 to l-1.
4. Inside the loop, it checks if the current character and the next character form a subtractive pair (e.g., IV, IX, XL, etc.). If so, it adds the corresponding value to the number variable and increments the loop index by 1.
5. If the loop index reaches l (indicating the last character of the input string), it sets the end flag to true.
6. If the current character and the next character do not form a subtractive pair, it adds the value of the current character to the number variable.
7. After the loop, if the end flag is false, it adds the value of the last character of the input string to the number variable.
8. Finally, the function returns the calculated number.

### Solution 2
#### RomanToInt2
Iterate over string from the end and check if current symbol is less than previous one. If it is, then subtract current symbol from result. Otherwise just add current symbol to result.

1. Check if validationRequired is true and if the input string s is a valid Roman numeral using the isValidRoman function. If it is not valid, return 0.
2. Create a map hm that maps each Roman numeral character to its corresponding value.
3. Initialize variables cur, total, and v to 0.
4. Iterate through the input string s from right to left.
5. Get the value of the current character v from the map hm.
6. If the current value cur is greater than v, subtract v from total.
7. Otherwise, add v to total and update cur to v.
8. Return the final value of total.


## benchmark result
```bash
goos: windows
goarch: amd64
pkg: github.com/slavakukuyev/leetcode-go/roman_to_int
cpu: 12th Gen Intel(R) Core(TM) i7-1255U

=== RUN   BenchmarkRomanToInt
BenchmarkRomanToInt
BenchmarkRomanToInt-12
  152348              8123 ns/op            9704 B/op         66 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/roman_to_int        1.345s

=== RUN   BenchmarkRomanToInt2
BenchmarkRomanToInt2
BenchmarkRomanToInt2-12
  140019              8150 ns/op            9705 B/op         66 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/roman_to_int        1.257s

=== RUN   BenchmarkRomanToIntWithoutValidation
BenchmarkRomanToIntWithoutValidation
BenchmarkRomanToIntWithoutValidation-12
15371136                76.77 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/roman_to_int        1.296s

=== RUN   BenchmarkRomanToInt2WithoutValidation
BenchmarkRomanToInt2WithoutValidation
BenchmarkRomanToInt2WithoutValidation-12
 5153563               228.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/roman_to_int        1.447s
```

