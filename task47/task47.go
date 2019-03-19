package task47

func permuteUnique(nums []int) [][]int {
	var result [][]int

	recursion(nums, 0, &result)

	return result
}

func recursion(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}

	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if !used[nums[i]] {
			used[nums[i]] = true

			swap(nums, i, start)
			recursion(nums, start+1, result)
			swap(nums, start, i)
		}
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
