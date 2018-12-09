package tasks

// 计数并说出
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prevStr := countAndSay(n - 1)
	s := ""
	for len(prevStr) != 0 {
		count := 1
		k := prevStr[0]
		i := 1
		for ; i < len(prevStr) && k == prevStr[i]; i++ {
			count++
		}
		prevStr = prevStr[i:]

		s += string('0'+count) + string(k)
	}

	return s
}
