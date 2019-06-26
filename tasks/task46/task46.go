package task46

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	revPermutes := [][]int{}
	prevPermutes := permute(nums[1:])
	for _, prevPermute := range prevPermutes {
		for i := 0; i <= len(prevPermute); i++ {
			t := append([]int{}, prevPermute[:i]...)
			t = append(t, nums[0])
			t = append(t, prevPermute[i:]...)
			revPermutes = append(revPermutes, t)
		}
	}

	return revPermutes
}
