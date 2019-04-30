package task49

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[[26]int][]string)

	for _, str := range strs {
		alphabet := [26]int{}

		for _, s := range str {
			alphabet[s-'a']++
		}

		groupMap[alphabet] = append(groupMap[alphabet], str)
	}

	result := make([][]string, 0, len(groupMap))

	for _, group := range groupMap {
		result = append(result, group)
	}

	return result
}
