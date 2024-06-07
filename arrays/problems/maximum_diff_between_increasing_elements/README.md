# Maximum Difference Between Increasing Elements

Given a 0-indexed integer array nums of size n, find the maximum difference between nums[i] and nums[j] (i.e., nums[j] - nums[i]), such that 0 <= i < j < n and nums[i] < nums[j].

Return the maximum difference. If no such i and j exists, return -1.

## Algorithm

1. initialize two variables
   1. maxDifference which is initialized to -1
   2. minElement which is initialized to the first element of the input array
2. loop through the input array from the second element
3. calculate the minElement which is min of current element or the minElement
4. calculate the maxDifference which is the max of the current maxDifference or minElement - current element
5. if the maxDifference is less than or equal to 0 after the loop, return -1
6. else, return maxDifference