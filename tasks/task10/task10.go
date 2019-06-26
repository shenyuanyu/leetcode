package task10

func isMatch(s string, p string) bool {
	ps := splitPattern([]byte(p))

	return subIsMatch(s, ps)
}

func subIsMatch(s string, ps [][]byte) bool {
	if len(s) != 0 && len(ps) == 0 {
		return false
	}
	if len(s) == 0 {
		for _, p := range ps {
			if !pattenEqual(s, p) {
				return false
			}
		}

		return true
	}

	isEqual := false
	for i := 0; i <= len(s); i++ {
		if pattenEqual(s[:i], ps[0]) {
			isEqual = subIsMatch(s[i:], ps[1:])
			if isEqual {
				return true
			}
		}
	}

	return false
}

func splitPattern(p []byte) [][]byte {
	if len(p) == 0 {
		return nil
	}
	if len(p) < 3 {
		return [][]byte{p}
	}

	ps := make([][]byte, 0, 10)

	t := []byte{}
	for i := 0; i < len(p); i++ {

		if i < len(p)-1 && p[i+1] == '*' {
			if len(t) != 0 {
				ps = append(ps, t)
			}
			ps = append(ps, []byte{p[i], p[i+1]})
			t = []byte{}
			i++
			continue
		}

		t = append(t, p[i])
	}

	if len(t) != 0 {
		ps = append(ps, t)
	}

	return ps
}

func pattenEqual(s string, p []byte) bool {
	if p[len(p)-1] == '*' { // 是正则
		if p[0] == '.' {
			return true
		}
		for i := 0; i < len(s); i++ {
			if s[i] != p[0] {
				return false
			}
		}
		return true
	}

	// 不是正则
	if len(s) != len(p) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if p[i] == '.' {
			continue
		}

		if s[i] != p[i] {
			return false
		}
	}

	return true
}
