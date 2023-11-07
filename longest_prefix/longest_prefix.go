package longestprefix

import (
	"bytes"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	if 1 > len(strs) || len(strs) > 200 {
		return ""
	}

	if len(strs[0]) < 1 || len(strs[0]) > 200 {
		return ""
	}

	latest := strs[0]
	var prefix bytes.Buffer
	for i := 1; i < len(strs); i++ {
		prefix.Reset()
		for j := 0; j < len(latest); j++ {
			if j < len(latest) && j < len(strs[i]) && latest[j] == strs[i][j] {
				prefix.WriteByte(strs[i][j])
			} else {
				break
			}
		}

		latest = prefix.String()
	}

	return latest
}

func LongestCommonPrefixLeetCode(strs []string) string {
	lenStr := len(strs)
	if lenStr == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < lenStr; j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

func LongestCommonPrefixCopilot(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
