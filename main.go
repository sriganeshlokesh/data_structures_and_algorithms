package main

import (
	"fmt"
	"sriganeshlokesh/data_structures/arrays/problems/best_time_to_buy_and_sell_stock"
	"sriganeshlokesh/data_structures/arrays/problems/maximum_diff_between_increasing_elements"
)

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maximum_diff_between_increasing_elements.MaxDiffIncreasingElements([]int{1, 5, 2, 10}))
}
