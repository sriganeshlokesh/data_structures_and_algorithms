package merge_intervals

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0]
	})

	mergedList := [][]int{intervals[0]}
	index := 0

	for i := 0; i < len(intervals); i++ {
		currentInterval := intervals[i]
		if mergedList[index][1] < currentInterval[0] {
			mergedList = append(mergedList, currentInterval)
			index++
		} else {
			mergedList[index][1] = max(mergedList[index][1], currentInterval[1])
		}
	}

	return mergedList
}
