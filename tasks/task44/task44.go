package task44

import "strings"

func isMatch(s string, p string) bool {
	patterns := splitPattern(p)

	if len(s) < stringSliceLen(patterns) {
		return false
	}
	return match(s, patterns)
}

// match ...
func match(s string, patterns []string) bool {
	if len(s) < stringSliceLen(patterns) { // 是否够长
		return false
	}

	if len(s) == 0 {
		return true
	}

	if len(patterns) == 0 {
		return false
	}

	if !isWildCardPattern(patterns[0]) { // 不是通配符
		if len(s) == len(patterns[0]) { // 长度相等
			if equal(s, patterns[0]) {
				for i := 1; i < len(patterns); i++ {
					if patterns[i] == "*" {
						continue
					}
					return false
				}
				return true
			}
			return false
		}

		// 长度不等
		if !equal(s[:len(patterns[0])], patterns[0]) {
			return false
		}
		// 递归
		return match(s[len(patterns[0]):], patterns[1:])
	}

	// 通配符
	if stringSliceLen(patterns) == 0 {
		return true
	}
	for i := 0; i < len(s)+1-len(patterns[1]); i++ {
		if equal(s[i:i+len(patterns[1])], patterns[1]) {
			if match(s[i+len(patterns[1]):], patterns[2:]) {
				return true
			}
		}
	}

	return false
}

// splitPattern 分离pattern
func splitPattern(p string) []string {
	var patterns []string
	var ptnBuilder strings.Builder
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			ptnBuilder.WriteByte(p[i])
			continue
		}

		if i == 0 {
			patterns = append(patterns, "*")
		}

		if ptnBuilder.Len() != 0 {
			patterns = append(patterns, ptnBuilder.String(), "*")
		}

		ptnBuilder.Reset()
	}

	if ptnBuilder.Len() != 0 {
		patterns = append(patterns, ptnBuilder.String())
	}
	return patterns
}

// equal 判断字符串与模式是否匹配
func equal(s string, p string) bool {
	if p == "*" {
		return true
	}
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

// isWildCardPattern 判断是否为模式
func isWildCardPattern(p string) bool {
	return p == "*"
}

// stringSliceLen 模式字符串切片中, 所有定长模式字符串之和
func stringSliceLen(sSlice []string) int {
	l := 0
	for _, s := range sSlice {
		if !isWildCardPattern(s) {
			l += len(s)
		}
	}
	return l
}
