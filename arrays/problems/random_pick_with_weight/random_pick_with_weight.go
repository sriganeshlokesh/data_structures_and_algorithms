package random_pick_with_weight

import "math/rand"

type Solution struct {
	pivots   []int
	maxValue int
}

func Constructor(w []int) Solution {
	pivots := make([]int, len(w)+1)
	for idx, weight := range w {
		pivots[idx+1] = pivots[idx] + weight
	}

	return Solution{pivots: pivots, maxValue: pivots[len(pivots)-1]}
}

func (s *Solution) PickIndex() int {
	random := rand.Intn(s.maxValue)

	start, end := 0, len(s.pivots)-1

	for start <= end {
		toCheck := (start + end) / 2
		if s.pivots[toCheck] > random {
			end = toCheck - 1
		} else {
			start = toCheck + 1
		}
	}

	return start - 1
}
