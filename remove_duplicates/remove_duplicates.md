# remove duplicates from a sorted array

## Problem
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

If all assertions pass, then your solution will be accepted.

## Examples
### Example 1:
```
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

### Example 2:
```
Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

## Constraints
```
1 <= nums.length <= 3 * 10**4
-10**4 <= nums[i] <= 10**4

nums is sorted in non-decreasing order.
```

## Flow
### RemoveDuplicates
```
Check the length of the input slice nums. If it is less than 3, return the length as there are no duplicates to remove.
Initialize a boolean variable max to keep track of whether the maximum number of duplicates (2) has been reached.
Initialize two pointers, left and pointer, to keep track of the current position in the slice.
Iterate through the slice starting from the second element (pointer = 1) until the end.
If the element at left is equal to the element at pointer and the maximum number of duplicates has not been reached (!max), increment left and set max to true.
If the element at left is not equal to the element at pointer, reset max to false and increment left.
Assign the element at pointer to the position at left in the slice.
Increment pointer to move to the next element.
Repeat steps 5-8 until all elements have been processed.
Return left + 1, which represents the length of the modified slice.
```

## Benchmarks
```
=== RUN   BenchmarkRemoveDuplicates
BenchmarkRemoveDuplicates
BenchmarkRemoveDuplicates-12
346219028                3.473 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/slavakukuyev/leetcode-go/remove_duplicates   10.430s
```


## Copilot
I wrote the explanation for a function to create Copilot version of the solution. Copilot generated for me the code I used for leetcode. I feel now like a Sage from  programmer's jokes memes :-)
