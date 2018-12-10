package tasks

// combinationSum 计算目标值
func combinationSum(candidates []int, target int) [][]int {
	rev := make([][]int, 0, 10)

	for i, candidate := range candidates {
		k := target - candidate
		if k > 0 {
			for _, t := range combinationSum(candidates[i:], k) {
				rev = append(rev, append(t, candidate))
			}
		} else if k == 0 {
			rev = append(rev, []int{candidate})
		}
	}

	return rev
}
