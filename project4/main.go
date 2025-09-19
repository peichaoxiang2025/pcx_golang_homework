package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // 假设第一个字符串为最长公共前缀
	for i := 1; i < len(strs); i++ {
		// 不断缩短前缀直到匹配当前字符串
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func main() {
	testCases := []string{ // 修正变量名为复数形式
		"flower",
		"flow",
		"flight",
	}

	result := longestCommonPrefix(testCases)
	fmt.Printf("Input: %v → Longest Common Prefix: %q\n", testCases, result)

	// 新增测试用例验证
	testCases2 := []string{
		"abcdef",
		"abcrgjkcdef",
		"abcdfgkiocdef",
	}
	result2 := longestCommonPrefix(testCases2)
	fmt.Printf("Input: %v → Longest Common Prefix: %q\n", testCases2, result2)

}
