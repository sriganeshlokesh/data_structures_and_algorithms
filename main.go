package main

import (
	"fmt"
	"sriganeshlokesh/data_structures/arrays/problems/buildings_with_an_ocean_view"
	"sriganeshlokesh/data_structures/arrays/problems/merge_intervals"
)

func main() {
	fmt.Println(buildings_with_an_ocean_view.FindBuildings([]int{1, 3, 2, 4}))
	fmt.Println(merge_intervals.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
