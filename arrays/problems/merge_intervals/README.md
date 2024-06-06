# Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

## Algorithm

1. sort the array based on the first element of each item in the array
2. initialize a variable to hold the merged intervals and append the first item of the sorted input array
3. loop through the sorted array from the second element
4. if the last element of the last item inside the merged intervals array is less than the first element of the current item, then merge the two intervals
   1. when merging, take the max of the last element of the two intervals and make that as the last element of the merged interval
5. else, just append the new interval to the merged intervals array
6. return merged intervals array