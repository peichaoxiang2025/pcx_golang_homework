package main

import (
	"fmt"
)

func isPalindrome(number int) bool {
	//特殊处理
	if number < 0 || (number%10 == 0 && number != 0) {
		return false
	}

	//number反转
	reversedHalfNumber := 0

	//将bumber翻转后存储在reversedHalfNumber中
	for number > reversedHalfNumber {
		reversedHalfNumber = reversedHalfNumber*10 + number%10
		number /= 10
	}

	return number == reversedHalfNumber || number == reversedHalfNumber/10
}

func main() {
	testCases := []int{121, -121, 10, 0, 12321}
	for _, num := range testCases {
		fmt.Printf("%d is palindrome: %v\n", num, isPalindrome(num))
	}
}
