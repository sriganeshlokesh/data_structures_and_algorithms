package main

import (
	"fmt"
	"sriganeshlokesh/data_structures/arrays/problems/intersection_of_two_arrays"
	"sriganeshlokesh/data_structures/arrays/problems/missing_number"
)

func main() {
	fmt.Println(intersection_of_two_arrays.Intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(missing_number.MissingNumber([]int{3, 0, 1}))
}
