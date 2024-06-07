# Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

## Algorithm

1. initialize two variables
   1. maximumProfit to 0
   2. minPrice to the first element of the input array
2. loop through the input array from the second element
3. if the current element is less than the minPrice, assign that value as minPrice
4. calculate the maximumProfit which is the max of current maximumProfit or the difference between the current element and the minPrice
5. return maximumProfit