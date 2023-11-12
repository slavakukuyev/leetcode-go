# zigzag string

## Description

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)
```text
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
```text
string convert(string text, int nRows);
```
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

## Functions
### Convert
```
The function first checks if numRows is equal to 1 or if the length of the string s is equal to 1. If either of these conditions is true, the function returns the original string s as there is no need to convert it.
The function initializes variables currentIndex to 0, stack as an empty slice of byte slices, ascending as true, and tempArray as a byte slice of length numRows.
The function iterates over each character in the string s using a for loop.
Inside the loop, if ascending is true, the current character is added to the tempArray at the currentIndex position, and currentIndex is incremented by 1.
If ascending is false, the current character is added to the tempArray at the currentIndex position only if currentIndex is not 0 or numRows-1. Otherwise, i is decremented by 1 to repeat the iteration with the same character.
After adding the character to tempArray, the function checks if currentIndex is equal to numRows or -1. If so, it appends tempArray to the stack, toggles the value of ascending, and creates a new tempArray with length numRows.
The function then checks if currentIndex is equal to numRows or -1 again and adjusts currentIndex accordingly.
After the loop, the tempArray is appended to the stack.
The function creates a new bytes.Buffer called str to store the final result.
It iterates over each row from 0 to numRows and for each row, iterates over each element in the stack.
If the element at the current row and column in the stack is not 0, it is added to the str buffer.
Finally, the function returns the string representation of the str buffer.
```

### ConvertCopilot
```
If numRows is equal to 1, return the input string s as it is.
Create an empty byte slice result to store the converted string.
Iterate over each row from 0 to numRows-1:
For each row i, iterate over the characters in the input string s starting from index i and incrementing by 2*numRows-2 in each iteration.
Append the character at index j to the result slice.
If i is not equal to 0 and not equal to numRows-1 and there is a character at index j+2*numRows-2-2*i in the input string s, append that character to the result slice as well.
Convert the result byte slice to a string and return it.
```

## Benchmark
```
=== RUN   BenchmarkConvert
BenchmarkConvert
BenchmarkConvert-12
 3233786               369.7 ns/op           462 B/op         13 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/convert_zigzag_string       1.609s


=== RUN   BenchmarkConvertCopilot
BenchmarkConvertCopilot
BenchmarkConvertCopilot-12
16628074                70.83 ns/op           40 B/op          3 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/convert_zigzag_string       1.278s
```
