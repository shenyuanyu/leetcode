package task44

import (
	"strings"
)

func isMatch(s string, p string) bool {
	p = pretreatment(p)
	if p == "*" {
		return true
	}
	if len(p) == 0 {
		return s == ""
	}

	patterns, startWithWildCard, endWithWildCard := extractPattern(p)
	if endWithWildCard { // end with *
		startPos := 0
		for i := 0; i < len(patterns); i++ {
			pos := stringsPatternIndex(s[startPos:], patterns[i])
			if pos == -1 {
				return false
			}
			if !startWithWildCard && startPos == 0 && pos != 0 {
				return false
			}

			startPos += pos + len(patterns[i])
		}
		return true
	}

	if startWithWildCard { // start with *
		endPos := len(s)
		for i := len(patterns) - 1; i >= 0; i-- {
			pos := stringsPatternLastIndex(s[:endPos], patterns[i])
			if pos == -1 {
				return false
			}
			if !endWithWildCard && endPos == len(s) && pos != len(s) {
				return false
			}

			endPos = pos - len(patterns[i])
		}
		return true
	}

	// not start with * and not end with *
	if len(patterns) == 1 {
		return equalPattern(s, patterns[0])
	}

	first := stringsPatternIndex(s, patterns[0])
	if first != 0 {
		return false
	}
	first += len(patterns[0])
	for i := 1; i < len(patterns)-1; i++ {
		pos := stringsPatternIndex(s[first:], patterns[i])
		if pos == -1 {
			return false
		}

		first += pos + len(patterns[i])
	}

	last := stringsPatternLastIndex(s[first:], patterns[len(patterns)-1])
	if last != len(s[first:]) {
		return false
	}

	return true
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

func extractPattern(p string) ([]string, bool, bool) {
	patterns := make([]string, 0, 10)

	startWithWildCard, endWithWildCard := false, false
	if len(p) != 0 {
		startWithWildCard = p[0] == '*'
		endWithWildCard = p[len(p)-1] == '*'
	}
	var str strings.Builder
	for _, c := range p {
		if c != '*' {
			str.WriteRune(c)
			continue
		}

		if str.Len() != 0 {
			patterns = append(patterns, str.String())
		}

		str.Reset()
	}

	if str.Len() != 0 {
		patterns = append(patterns, str.String())
	}

	return patterns, startWithWildCard, endWithWildCard
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

func stringsPatternIndex(s string, p string) int {
	for i := 0; i <= len(s)-len(p); i++ {
		if equalPattern(s[i:i+len(p)], p) {
			return i
		}
	}
	return -1
}

func stringsPatternLastIndex(s string, p string) int {
	for i := len(s); i >= len(p); i-- {
		if equalPattern(s[i-len(p):i], p) {
			return i
		}
	}
	return -1
}
