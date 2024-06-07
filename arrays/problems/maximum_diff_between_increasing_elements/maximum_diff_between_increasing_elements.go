package maximum_diff_between_increasing_elements

func MaxDiffIncreasingElements(nums []int) int {
	maxDifference := -1
	minElement := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < minElement {
			minElement = nums[i]
		}

		maxDifference = max(maxDifference, nums[i]-minElement)
	}

	if maxDifference <= 0 {
		return -1
	}

	return maxDifference
}
