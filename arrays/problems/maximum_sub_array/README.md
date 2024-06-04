# Maximum Sub Array

Given an integer array nums, find the subarray
with the largest sum, and return its sum.

## Algorithm

1. initialize two variables maxSum and currentSum which will represent the maximum sum and the current maximum sum
   1. initialize them to the first element of the input array
2. loop through the input array from the second element
3. if the currentSum is less than 0, we just assign the currentSum to the current element of the array
4. else, we add the current element of the array to the currentSum
5. we then calculate the max of maxSum and currentSum and assign it to maxSum variable
6. return maxSum