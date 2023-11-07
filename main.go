package main

import (
	"fmt"

	longestprefix "github.com/slavakukuyev/leetcode-go/longest_prefix"
	romantoint "github.com/slavakukuyev/leetcode-go/roman_to_int"
)

func main() {
	fmt.Println(romantoint.RomanToInt("MCMXCIV", false))  //1994
	fmt.Println(romantoint.RomanToInt2("MCMXCIV", false)) //1994

	fmt.Println(longestprefix.LongestCommonPrefix([]string{"flower", "flow", "flight"}))                   //fl
	fmt.Println(longestprefix.LongestCommonPrefixLeetCode([]string{"flower", "flow", "flight", "cardfg"})) //fl
	fmt.Println(longestprefix.LongestCommonPrefixCopilot([]string{"flower", "flow", "flight", "cardfg"}))  //fl
}
