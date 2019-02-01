package task43

import (
	"strconv"
	"strings"
)

// multiply ...
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	pos := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			pos[i+j+1] += int((num1[i] - '0') * (num2[j] - '0'))
		}
	}

	carry := 0
	for i := len(pos) - 1; i >= 0; i-- {
		pos[i] += carry
		carry = pos[i] / 10
		pos[i] %= 10
	}

	i := 0
	if pos[i] == 0 {
		i = 1
	}
	var sb strings.Builder
	for ; i < len(pos); i++ {
		sb.WriteString(strconv.Itoa(pos[i]))
	}

	return sb.String()
}
