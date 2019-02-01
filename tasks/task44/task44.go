package task44

import "strings"

func isMatch(s string, p string) bool {
	if len(s) < len(p) {
		return false
	}

	patterns := splitPattern(p)

}

func match(s string, patterns []string) bool {
	if len(s) < stringSliceLen(patterns) {
		return false
	}

	for _, pattern := range patterns {
		if !isWildCardPattern(pattern) {
			l := len(pattern)
			if s[:l] == pattern {
				return match(s[l:], patterns)
			}
		}

	}
}

func splitPattern(p string) []string {
	var patterns []string
	var ptnBuilder strings.Builder
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			ptnBuilder.WriteByte(p[i])
			continue
		}

		patterns = append(patterns, ptnBuilder.String(), "*")
		ptnBuilder.Reset()
	}

	if ptnBuilder.Len() != 0 {
		patterns = append(patterns, ptnBuilder.String())
	}
	return patterns
}

func equal(s string, p string) bool {
	if p == "*" {
		return true
	}
	if s == p {
		return true
	}
	if p[len(p)-1] == '?' {
		return s[:len(s)-1] == p[:len(p)-1]
	}
	return false
}

func isWildCardPattern(p string) bool {
	return p == "*"
}

func stringSliceLen(sSlice []string) int {
	l := 0
	for _, s := range sSlice {
		l += len(s)
	}
	return l
}
