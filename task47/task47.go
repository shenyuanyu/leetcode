package task47

import (
	"strconv"
	"strings"
)

func permuteUnique(nums []int) [][]int {
	resultMap := make(map[string][]int)

	recursion(nums, 0, resultMap)

	result := make([][]int, 0, len(resultMap))
	for _, val := range resultMap {
		result = append(result, val)
	}

	return result
}

func recursion(nums []int, start int, resultMap map[string][]int) {
	if start == len(nums)-1 {
		key := intSliceToString(nums)
		if _, ok := resultMap[key]; ok {
			return
		}

		tmp := make([]int, len(nums))
		copy(tmp, nums)
		resultMap[key] = tmp
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[start] {
			continue
		}

		swap(nums, i, start)
		recursion(nums, start+1, resultMap)
		swap(nums, start, i)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func intSliceToString(nums []int) string {
	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}

	return sb.String()
}
