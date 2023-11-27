# max area

## Problem Statement
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

## Examples
### Example 1:
```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```

### Example 2:
```
Input: height = [1,1]
Output: 1
```

## Flow
### MaxArea
```text
Initialize variables max, left, right, w, and h to 0.
Iterate while left is less than right:
Calculate the width w as the difference between right and left.
If the height at left is less than the height at right:
Set h to the height at left.
Increment left by 1.
Else:
Set h to the height at right.
Decrement right by 1.
Calculate the area w * h.
If the calculated area is greater than max, update max with the new area.
Return the maximum area max.
```

### MaxAreaCopilot
```text
Initialize variables max, left, right, w, and h to 0.
Iterate while left is less than right:
Calculate the width w as the difference between right and left.
If the height at left is less than the height at right:
Set h to the height at left.
Increment left by 1.
Else:
Set h to the height at right.
Decrement right by 1.
Calculate the area w * h.
If the calculated area is greater than max, update max with the new area.
Return the maximum area max.
```

## Benchmark
```
=== RUN   BenchmarkMaxArea
BenchmarkMaxArea
BenchmarkMaxArea-12
287115682                4.208 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/max_area    10.689s

=== RUN   BenchmarkMaxAreaCopilot
BenchmarkMaxAreaCopilot
BenchmarkMaxAreaCopilot-12
266272346                4.410 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/max_area    1.706s
```