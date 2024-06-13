# Missing Number

## Problem Statement
Given an array nums containing n distinct numbers taken from the range 0 to n, find the one number that is missing from the array.

## Solution
The goal is to identify the missing number in the array. We can solve this problem efficiently using a mathematical approach based on the properties of arithmetic series.

### Approach
* We use the fact that the sum of the first n natural numbers is given by the formula n * (n + 1) / 2 
* By calculating the expected sum of numbers from 0 to n and subtracting the sum of the elements in the array, we can find the missing number.

### Detailed Steps

* Calculate the length of the input array nums and store it in size. 
* Initialize two variables, current and total, to 0. current will store the sum of elements in nums, and total will store the sum of numbers from 0 to n. 
* Iterate from 0 to size inclusive (size + 1 iterations). 
* For each iteration:
  * Add the current index i to total to accumulate the sum of all numbers from 0 to n. 
  * If i is less than size, add the element nums[i] to current to accumulate the sum of the elements in nums. 
  * Finding the Missing Number:
    * The missing number is the difference between total and current.

### Complexity
**Time Complexity**: O(n), where n is the number of elements in nums. The loop runs size + 1 times, which is effectively linear.

**Space Complexity**: O(1), as no additional space is used that grows with the input size.