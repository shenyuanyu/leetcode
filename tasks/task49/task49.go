package task49

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[[26]byte]int)

	result := make([][]string, 0)
	for _, str := range strs {
		alphabet := [26]byte{}

		for _, s := range str {
			alphabet[s-'a']++
		}

		if index, ok := groupMap[alphabet]; ok {
			result[index] = append(result[index], str)
		} else {
			groupMap[alphabet] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}
