package maximum_sub_array

func MaximumSubArray(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]

	for _, num := range nums[1:] {
		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum += num
		}

		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
