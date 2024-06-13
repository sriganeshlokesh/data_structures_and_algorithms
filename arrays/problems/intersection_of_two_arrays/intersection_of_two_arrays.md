# Intersection of Two Arrays

## Problem Statement
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

## Solution
The goal is to find the intersection of two arrays, nums1 and nums2, meaning the set of elements that appear in both arrays.

### Approach
We can solve this problem efficiently by leveraging a counting mechanism using an array to keep track of elements in nums1, and then iterating through nums2 to find common elements.

### Detailed Steps
#### Initialization:

* Initialize an array seen of size 1000 with all elements set to 0. This array will be used to count occurrences of elements in nums1.
Counting Occurrences in nums1:
* Iterate over each element in nums1 and increment the corresponding index in the seen array. This effectively counts how many times each element appears in nums1.
Finding Common Elements:
* Initialize an empty slice res to store the result. 
* Iterate over each element in nums2. For each element, check if the corresponding count in the seen array is greater than 0 (indicating that the element is present in nums1). 
* If the element from nums2 is found in nums1, append this element to the result slice res and set the count in seen to 0 to avoid duplicate entries in the result.

#### Return the Result:

Finally, return the slice res which contains the intersection of nums1 and nums2.

### Complexity
**Time Complexity**: O(n + m), where n is the length of nums1 and m is the length of nums2. This is because we iterate through both arrays once.

**Space Complexity**: O(1) if the range of elements is fixed and small (0 to 999 in this case).