package utils

import (
	"strings"
)

func ConcatenateStrings(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
