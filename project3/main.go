package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	//定义map存储括号
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	//定义切片
	stack := []rune{}

	for _, char := range s {
		// 如果是右括号
		if matching, exists := bracketMap[char]; exists {
			// 栈为空或栈顶元素不匹配则无效
			if len(stack) == 0 || stack[len(stack)-1] != matching {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
	}

	for _, tc := range testCases {
		result := isValid(tc.input)
		fmt.Printf("Input: %q, Expected: %v, Got: %v\n", tc.input, tc.expected, result)
	}
}
