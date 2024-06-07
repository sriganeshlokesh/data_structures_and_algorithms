# Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.


## Algorithm

1. initialize a map with key as the number and the value is the index of the number in the input
2. loop through the array of integers
   1. check if there is a number in the map that is the result of `target - current` number
      1. if true, return the index of that number and the current index
   2. else, add the number as key and its index to the map
3. if no match is found, return an empty list