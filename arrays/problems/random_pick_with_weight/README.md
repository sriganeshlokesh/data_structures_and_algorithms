# Random Pick With Weight

## Intuition
When dealing with weighted random selection, we aim to pick an index such that the probability of picking any index is proportional to its weight.

Consider weights [10, 200, 30, 40]. The total weight is 280. The probability of picking each weight should correspond to the ratio of the weight to the total weight.

To visualize, imagine the weights as intervals on a number line from 0 to 280:

Weight 10 spans 0 to 10 (exclusive, i.e., [0, 10)).
Weight 200 spans 10 to 210 (exclusive, i.e., [10, 210)).
Weight 30 spans 210 to 240 (exclusive, i.e., [210, 240)).
Weight 40 spans 240 to 280 (exclusive, i.e., [240, 280)).
When a random number is picked from 0 to 280 (exclusive), the range where the number falls determines the chosen index:

0-9: Index 0
10-209: Index 1
210-239: Index 2
240-279: Index 3

## Algorithm
**Constructor**:

1. We build the pivots array, which marks the cumulative sums of weights. This helps in mapping a random number to its corresponding weight interval.
pivots[i] will represent the end of the interval for the i-1 weight.
PickIndex Function:
2. Generate a random number within the range 0 to maxValue (exclusive).
Use binary search to find the correct interval in which this random number falls.

**Initialize pivots**:

pivots := make([]int, len(w) + 1) creates an array of size len(w) + 1.
Fill pivots array:
For each weight, add it to the cumulative sum in the pivots array.
pivots[idx + 1] = pivots[idx] + value ensures each position in pivots contains the sum of weights up to that index.

**Return the Solution object**:
Solution{pivots: pivots, maxValue: pivots[len(pivots) - 1]} initializes the Solution struct with pivots and the maximum value (total sum of weights).

**Generate a random number**:

random := rand.Intn(this.maxValue) picks a random number in the range [0, this.maxValue).

**Binary Search**:

1. start and end initialize the search range.
2. In each iteration:
   1. Calculate toCheck as the midpoint: (start + end) / 2.
   If this.pivots[toCheck] > random, the number lies in the left half of the search range; adjust end. 
   2. Otherwise, adjust start to search the right half. 
   3. Continue until start surpasses end. 
   4. Return the index

return start - 1 gives the correct weight index corresponding to the interval containing the random number.

### Complexity

#### Time Complexity:

**Constructor**: O(n) to build the pivots array.

**PickIndex**: O(log n) due to the binary search within the pivots array.

#### Space Complexity:

**Constructor**: O(n) for storing the pivots array.

**PickIndex**: O(1) as it uses constant extra space for variables.
