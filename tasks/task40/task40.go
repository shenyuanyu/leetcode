package task40

import (
	"sort"
)

// 求和等于target的候选值, 候选值可能重复
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return subCombinationSum2(candidates, target)
}

func subCombinationSum2(candidates []int, target int) [][]int {
	rev := make([][]int, 0, 10)
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		k := target - candidates[i]
		if k > 0 {
			ts := combinationSum2(candidates[i+1:], k)
			for _, t := range ts {
				rev = append(rev, append(t, candidates[i]))
			}
		} else if k == 0 {
			rev = append(rev, []int{candidates[i]})
		}
	}

	return rev
}
