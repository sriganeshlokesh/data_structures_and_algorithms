package two_sum

func TwoSum(nums []int, target int) []int {
	pairs := make(map[int]int, len(nums))

	for i, num := range nums {
		if _, ok := pairs[target-num]; ok {
			return []int{pairs[target-num], i}
		}

		pairs[num] = i
	}

	return []int{}
}
