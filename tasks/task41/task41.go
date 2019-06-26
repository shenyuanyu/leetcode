package task41

func firstMissingPositive(nums []int) int {
	miss := 1
	hash := make(map[int]struct{})
	for _, num := range nums {
		hash[num] = struct{}{}

		for {
			if _, ok := hash[miss]; ok {
				miss++
			} else {
				break
			}
		}
	}

	return miss
}
