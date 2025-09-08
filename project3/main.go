package main

import (
	"fmt"
)

func isValid(s string) bool {
	//定义map存储括号
	bracketMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	//定义切片
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		//如果当前字符是左括号，将其压入栈中
		if value, exists := bracketMap[c]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

func main() {
	testCases := []int{121, -121, 10, 0, 12321}
	for _, num := range testCases {
		fmt.Printf("%d is palindrome: %v\n", num, isPalindrome(num))
	}
}
