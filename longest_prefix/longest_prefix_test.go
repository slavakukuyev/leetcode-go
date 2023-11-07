package longestprefix

import "testing"

func BenchmarkLongestPrefix(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix(strs)
	}
}

func BenchmarkLongestPrefixLeetCode(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefixLeetCode(strs)
	}
}

func BenchmarkLongestCommonPrefixCopilot(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefixCopilot(strs)
	}
}
