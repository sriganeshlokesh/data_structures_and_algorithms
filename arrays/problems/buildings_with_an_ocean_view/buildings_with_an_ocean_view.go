package buildings_with_an_ocean_view

import "slices"

func FindBuildings(heights []int) []int {
	visibleBuildings := make([]int, 0)
	maxHeight := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			visibleBuildings = append(visibleBuildings, i)
			maxHeight = heights[i]
		}
	}

	slices.Sort(visibleBuildings)

	return visibleBuildings
}
