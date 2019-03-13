package task45

func jump(nums []int) int {
	step := 0
	for i := 0; i < len(nums)-1; {
		i = nextPosition(nums, i)
		step++
	}

	return step
}

func nextPosition(nums []int, nowPos int) int {
	stepRange := nums[nowPos]

	if nowPos+stepRange >= len(nums)-1 {
		return len(nums) - 1
	}

	maxPos, targetPos := 0, 0
	for i := nowPos + 1; i <= nowPos+stepRange; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
			targetPos = i
		}
	}

	return targetPos
}
