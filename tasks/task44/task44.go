package task44

import "strings"

func isMatch(s string, p string) bool {
	p = pretreatment(p)
	patterns := splitPattern(p)

	return subIsMatch(s, patterns)
}

func subIsMatch(s string, patterns []string) bool {
	if len(patterns) == 0 {
		return len(s) == 0
	}

	if patterns[0] != "*" {
		if len(s) < len(patterns[0]) || !equalPattern(s[:len(patterns[0])], patterns[0]) {
			return false
		}
		return subIsMatch(s[len(patterns[0]):], patterns[1:])
	}

	if len(patterns) == 1 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if subIsMatch(s[i:], patterns[1:]) {
			return true
		}
	}

	return false
}

func pretreatment(p string) string {
	var str strings.Builder
	for i := 0; i < len(p); {
		str.WriteByte(p[i])
		if p[i] != '*' {
			i++
			continue
		}

		i++
		for i < len(p) && p[i] == '*' {
			i++
		}
	}

	return str.String()
}

func splitPattern(p string) []string {
	patterns := make([]string, 0, 10)
	var str strings.Builder
	for _, c := range p {
		if c != '*' {
			str.WriteRune(c)
			continue
		}

		if str.Len() != 0 {
			patterns = append(patterns, str.String(), string(c))
		} else {
			patterns = append(patterns, string(c))
		}

		str.Reset()
	}

	if str.Len() != 0 {
		patterns = append(patterns, str.String())
	}

	return patterns
}

func equalPattern(s string, p string) bool {
	if p != "*" {
		if len(s) != len(p) {
			return false
		}

		for i := 0; i < len(s); i++ {
			if s[i] != p[i] && p[i] != '?' {
				return false
			}
		}
		return true
	}

	return true
}
