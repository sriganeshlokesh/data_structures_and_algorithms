# Buildings With an Ocean View

There are n buildings in a line. You are given an integer array heights of size n that represents the heights of the buildings in the line.

The ocean is to the right of the buildings. A building has an ocean view if the building can see the ocean without obstructions. Formally, a building has an ocean view if all the buildings to its right have a smaller height.

Return a list of indices (0-indexed) of buildings that have an ocean view, sorted in increasing order.

## Algorithm

1. initialize two variables
   1. maximumHeight which indicates the current maximum height building from the input array
      1. this is initialized to -1 
   2. visibleBuildings which is an array of integers that will hold the ocean view building index
2. loop through the input array from the end as the last building is right next to the ocean
3. if the current building height is greater than the max height, we append the index to the visibleBuilding array and assign max height to this new building height
4. finally, since we started from the end and the problem asks us to return the index in sorted increasing order, we sort the array
5. return the visibleBuildings array