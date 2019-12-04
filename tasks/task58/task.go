package task58

func lengthOfLastWord(s string) int {
	// remove blank
	l := len(s) - 1
	for l >= 0 && s[l] == ' ' {
		l--
	}

	t := l
	for t >= 0 && s[t] != ' ' {
		t--
	}
	return l - t
}
