package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	// 从末位开始遍历
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 { // 优化：提前判断是否需要进位
			digits[i]++
			return digits
		}
		digits[i] = 0 // 进位后置零
	}

	// 如果所有位都是9（如[9,9,9]），需要扩展数组
	return append([]int{1}, digits...)
}

func main() {
	// 测试案例
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{0},
		{2, 0, 9},
		{1, 9, 9},
	}

	// 执行测试
	for _, tc := range testCases {
		// 复制输入以避免修改原数组
		input := make([]int, len(tc))
		copy(input, tc)

		result := plusOne(tc)
		fmt.Printf("输入: %v, 结果: %v\n", input, result)
	}
}
