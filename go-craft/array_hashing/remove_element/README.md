# Remove Element

## Problem Description

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in-place. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

## Approach

- **Two Pointers**: Use a slow pointer to track the position where the next valid element should be placed, and a fast pointer to iterate through the array.
- **Time Complexity**: O(n)
- **Space Complexity**: O(1)

## Examples

```go
// Example 1:
nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]

// Example 2:
nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,3,0,4,_,_,_]
```

## Test Cases

Run tests with: `go test -v`
Run benchmarks with: `go test -bench=.`
