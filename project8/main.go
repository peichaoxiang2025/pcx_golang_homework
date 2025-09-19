package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 哈希表存储数值到索引的映射
	for i, num := range nums {
		complement := target - num // 计算补数
		if j, exists := numMap[complement]; exists {
			return []int{j, i} // 找到解立即返回
		}
		numMap[num] = i // 存储当前数值和索引
	}
	return nil // 无解时返回nil（根据题意可改为抛出异常）
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4}, 10, nil},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("Input: %v Target: %d → Output: %v (Expected: %v)\n",
			tc.nums, tc.target, result, tc.expect)
	}
}
