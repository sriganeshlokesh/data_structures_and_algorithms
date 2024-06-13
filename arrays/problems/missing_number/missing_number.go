package missing_number

func MissingNumber(nums []int) int {
	size := len(nums)
	var current, total int

	for i := 0; i < size+1; i++ {
		total += i

		if i < size {
			current += nums[i]
		}
	}

	return total - current
}
