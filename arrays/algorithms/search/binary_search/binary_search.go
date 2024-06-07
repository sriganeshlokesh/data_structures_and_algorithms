package binary_search

func IterativeSearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func RecursiveSearch(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return RecursiveSearch(nums, start, mid-1, target)
	} else {
		return RecursiveSearch(nums, mid+1, end, target)
	}
}
