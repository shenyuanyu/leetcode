package task55

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	farthest := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && farthest == i && i != len(nums)-1 {
			return false
		}

		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
	}

	return true
}
