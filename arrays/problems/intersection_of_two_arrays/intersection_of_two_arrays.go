package intersection_of_two_arrays

func Intersection(nums1 []int, nums2 []int) []int {
	seen := make([]int, 1001)
	for i := 0; i < len(nums1); i++ {
		seen[nums1[i]]++
	}

	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if seen[nums2[i]] > 0 {
			res = append(res, nums2[i])
			seen[nums2[i]] = 0
		}
	}

	return res
}
