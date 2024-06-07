# Binary Search

Binary search is an efficient algorithm for finding an item from a sorted list of items. It works by repeatedly dividing in half the portion of the list that could contain the item until you've narrowed down the possible locations to just one.

## Detailed Explanation of Binary Search

The binary search algorithm works on the principle of divide and conquer. It operates on a sorted array and uses the following steps:

**Initialization**: Start with two pointers, start and end, which initially point to the first and last indices of the array.

**Middle Calculation**: Calculate the middle index (mid) of the current segment.

**Comparison**: Compare the target value with the value at the mid index:
If the target value is equal to the middle value, you've found the target.
If the target value is less than the middle value, narrow the search to the left half of the array.
If the target value is greater than the middle value, narrow the search to the right half of the array.

**Repeat**: Continue the process on the new segment until the segment is empty (i.e., start exceeds end).

### Complexity
**Time Complexity**:
𝑂
(
log
⁡
𝑛
)
O(logn) where
𝑛
n is the number of elements in the array. Each step effectively halves the search space.

**Space Complexity**:
𝑂
(
1
)
O(1) for the iterative version,
𝑂
(
log
⁡
𝑛
)
O(logn) for the recursive version due to the call stack.